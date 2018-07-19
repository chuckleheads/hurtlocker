#!/usr/bin/env ruby

require 'delegate'
require 'optparse'
require 'tsort'
require 'tempfile'

# Dependency tracker.
class Sortable < SimpleDelegator
  include TSort

  def add(ident, deps = [])
    __getobj__[ident] = deps
  end

  def tsort_each_node(&block)
    __getobj__.each_key(&block)
  end

  def tsort_each_child(node, &block)
    __getobj__.fetch(node).each(&block)
  end
end

def dependents(pkg_ident, all_deps)
  dependents = []
  all_deps.each do |ident, deps|
    if deps.include?(pkg_ident)
      dependents << ident
    end
  end
  dependents
end

def process(pkg_ident, all_deps)
  results = dependents(pkg_ident, all_deps)
  return if results.empty?

  @rebuilds.concat(results)
  results.each { |pkg_ident| process(pkg_ident, all_deps) }
end


bash_prog = Tempfile.new('print_deps.sh')
bash_prog.write(<<'EOF')
#!/bin/bash
set -e
STUDIO_TYPE=stage1
FIRST_PASS=true
PLAN_CONTEXT=$(pwd)/$(dirname $1)

cd $(dirname $1)
source $(basename $1)
echo "${pkg_origin}/${pkg_name}"
echo "${pkg_build_deps[*]} ${pkg_deps[*]}"
exit 0
EOF
bash_prog.close

all_deps = Sortable.new({})
ident_to_plan = {}

ARGF.each_line do |file|
  # Load the plan by sourcing it with a small Bash program
  raw = `bash #{bash_prog.path} #{file}`.chomp
  # Parse out the package identifier and the dependencies from the Bash program
  ident, _, deps_str = raw.partition(/\n/)
  puts "Ident: #{ident}: Deps: #{deps_str}"
  # Add the package ident to a key in an "all deps" hash with the value being
  # an array of dependency package identifiers, but drop ourself so we
  # don't an infinitely recursive result set. This would happen if we
  # depend on a previous version of ourself, for example, `go`
  # requires `go` to build a newer `go`
  all_deps.add(ident, deps_str.split(' ')
    .map { |d| d.split('/').first(2).join('/') }
    .reject { |i| i == ident })
  # Add the pacakge ident to a key in a "plan" hash with the value being
  # the path to the directory containing the relevant `plan.sh`
  ident_to_plan[ident] = \
    file.chomp.sub(%r{/habitat/plan.sh$}, '').sub(%r{/plan.sh}, '')
end

# Iterate through every dependency array entry and add an entry into the "all
# deps" hash if it does not yet exist. This is done to ensure the topological
# sort will continue to completion.
all_deps.keys.each do |ident|
  all_deps[ident].each do |dep|
    all_deps.add(dep) unless all_deps.key?(dep)
  end
end

puts "{\n"
puts "\tset {\n"
all_deps.keys.each do |ident|
  puts "\t\t_:#{ident.split('/')[1]} <ident> \"#{ident}\" .\n"
  all_deps[ident].each do |dep|
    puts "\t\t_:#{ident.split('/')[1]} <dependency> _:#{dep.split('/')[1]} .\n"
  end
end
puts "\t}\n"
puts "}\n"

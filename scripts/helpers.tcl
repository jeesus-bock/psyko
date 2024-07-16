# helper functions for the TCL scripts.

# brutally print out a dict as JSON.
proc dictToJson dic {
  puts "täs"
  puts $dic
  set ret "{ "
  dict for {k v} $dic {append ret "\"$k\":\"$v\","}
  set ret [string trimright $ret ,]
  append ret " } "
  return $ret
}


puts "first.tcl begin"
puts $data
set combined "{ "
dict for {k v} $data {append combined "\"$k\":\"$v\","}
set combined [string trimright $combined ,]
append combined " } "
return $combined


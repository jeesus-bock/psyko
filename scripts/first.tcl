# First demonstrative TCL script. Takes the data dict and returns it formatted as JSON.
puts "first.tcl begin"
puts $data
return [dictToJson $data]



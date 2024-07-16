# Playing around with raandom string elements and whatnot :)
proc randelem {list} {
    lindex $list [expr {int(rand()*[llength $list])}]
}
set x [randelem {"Karate" "Cooking" "Karaoke"}]
return $x

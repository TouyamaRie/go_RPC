module main

go 1.19

require geerpc v0.0.0
require geerpc/codec v0.0.0
require geerpc/xclient v0.0.0

replace geerpc => ../
replace geerpc/codec => ../codec
replace geerpc/xclient => ../xclient

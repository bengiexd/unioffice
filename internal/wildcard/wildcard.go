//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package wildcard ;func Index (pattern ,name string )(_ag int ){if pattern ==""||pattern =="\u002a"{return 0;};_ecd :=make ([]rune ,0,len (name ));_fc :=make ([]rune ,0,len (pattern ));for _ ,_fg :=range name {_ecd =append (_ecd ,_fg );};for _ ,_gadf :=range pattern {_fc =append (_fc ,_gadf );};return _ecb (_ecd ,_fc ,0);};func Match (pattern ,name string )(_a bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_cf :=make ([]rune ,0,len (name ));_eb :=make ([]rune ,0,len (pattern ));for _ ,_de :=range name {_cf =append (_cf ,_de );};for _ ,_b :=range pattern {_eb =append (_eb ,_b );};_cac :=false ;return _ee (_cf ,_eb ,_cac );};func _ee (_ba ,_ec []rune ,_ga bool )bool {for len (_ec )> 0{switch _ec [0]{default:if len (_ba )==0||_ba [0]!=_ec [0]{return false ;};case '?':if len (_ba )==0&&!_ga {return false ;};case '*':return _ee (_ba ,_ec [1:],_ga )||(len (_ba )> 0&&_ee (_ba [1:],_ec ,_ga ));};_ba =_ba [1:];_ec =_ec [1:];};return len (_ba )==0&&len (_ec )==0;};func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_c :=make ([]rune ,0,len (name ));_f :=make ([]rune ,0,len (pattern ));for _ ,_e :=range name {_c =append (_c ,_e );};for _ ,_ca :=range pattern {_f =append (_f ,_ca );};_g :=true ;return _ee (_c ,_f ,_g );};func _ecb (_gd ,_eeg []rune ,_eac int )int {for len (_eeg )> 0{switch _eeg [0]{default:if len (_gd )==0{return -1;};if _gd [0]!=_eeg [0]{return _ecb (_gd [1:],_eeg ,_eac +1);};case '?':if len (_gd )==0{return -1;};case '*':if len (_gd )==0{return -1;};_aa :=_ecb (_gd ,_eeg [1:],_eac );if _aa !=-1{return _eac ;}else {_aa =_ecb (_gd [1:],_eeg ,_eac );if _aa !=-1{return _eac ;}else {return -1;};};};_gd =_gd [1:];_eeg =_eeg [1:];};return _eac ;};
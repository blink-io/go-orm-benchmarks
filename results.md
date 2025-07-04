# Results

- orm-benchmark (with no flags)
```
Reports:

Insert
pgx:		3777	330568 ns/op	280 B/op	8 allocs/op
sq:		3632	338404 ns/op	9704 B/op	91 allocs/op
pg:		3234	356373 ns/op	1141 B/op	10 allocs/op
gorp:		4014	360562 ns/op	1735 B/op	33 allocs/op
bun:		3382	383468 ns/op	5028 B/op	13 allocs/op
pgx_pool:	2973	394602 ns/op	297 B/op	8 allocs/op
sqlboiler:	3492	400088 ns/op	1582 B/op	33 allocs/op
raw:		2751	410664 ns/op	704 B/op	13 allocs/op
beego:		2686	411289 ns/op	2400 B/op	56 allocs/op
jet:		2823	420187 ns/op	3528 B/op	98 allocs/op
ent:		3373	428089 ns/op	4150 B/op	97 allocs/op
reform:		3013	429122 ns/op	1775 B/op	51 allocs/op
gorm:		2716	430875 ns/op	7400 B/op	106 allocs/op
sqlc:		3127	431754 ns/op	280 B/op	8 allocs/op
gorm_prep:	2822	432913 ns/op	5385 B/op	66 allocs/op
dbr:		2866	502957 ns/op	2624 B/op	57 allocs/op
zorm:		2143	529382 ns/op	5080 B/op	95 allocs/op
gen:		2480	554548 ns/op	10617 B/op	134 allocs/op
godb:		2146	566107 ns/op	4409 B/op	99 allocs/op
sqlx:		1700	596612 ns/op	856 B/op	19 allocs/op
upper:		1818	598354 ns/op	5912 B/op	125 allocs/op
rel:		1768	605612 ns/op	2638 B/op	45 allocs/op
xorm:		2143	657057 ns/op	3120 B/op	87 allocs/op
pop:		1238	840256 ns/op	9318 B/op	212 allocs/op

InsertMulti
pgx:		1593	791043 ns/op	47615 B/op	38 allocs/op
pgx_pool:	1644	867685 ns/op	47658 B/op	38 allocs/op
sqlc:		1262	911923 ns/op	66657 B/op	639 allocs/op
sq:		1057	981090 ns/op	237159 B/op	1697 allocs/op
raw:		1256	996025 ns/op	187153 B/op	930 allocs/op
beego:		1128	1091401 ns/op	177724 B/op	2744 allocs/op
gorm_prep:	1094	1200142 ns/op	254592 B/op	1891 allocs/op
reform:		1017	1231801 ns/op	462189 B/op	2746 allocs/op
ent:		992	1292654 ns/op	396544 B/op	4597 allocs/op
pg:		860	1436381 ns/op	3315 B/op	112 allocs/op
jet:		862	1444461 ns/op	332733 B/op	5793 allocs/op
sqlx:		824	1497702 ns/op	170749 B/op	1550 allocs/op
bun:		705	1559377 ns/op	42638 B/op	219 allocs/op
gorm:		802	1560268 ns/op	276375 B/op	5231 allocs/op
gen:		804	1577989 ns/op	289954 B/op	5357 allocs/op
zorm:		790	1619524 ns/op	206519 B/op	2107 allocs/op
upper:		645	1730254 ns/op	328155 B/op	4204 allocs/op
godb:		619	1812262 ns/op	247937 B/op	4294 allocs/op
rel:		716	1846271 ns/op	312569 B/op	3265 allocs/op
xorm:		660	1896673 ns/op	258942 B/op	5518 allocs/op
gorp:		bulk-insert is not supported
sqlboiler:	bulk-insert is not supported
pop:		bulk-insert is not supported
dbr:		bulk-insert is not supported

Update
raw:		8305	149167 ns/op	750 B/op	13 allocs/op
sqlc:		8173	152389 ns/op	288 B/op	8 allocs/op
sqlx:		3547	329907 ns/op	872 B/op	20 allocs/op
sq:		3282	355395 ns/op	7343 B/op	81 allocs/op
pgx:		3277	371353 ns/op	288 B/op	8 allocs/op
pg:		2896	392764 ns/op	768 B/op	9 allocs/op
gorp:		2798	401618 ns/op	1133 B/op	23 allocs/op
sqlboiler:	3156	408959 ns/op	900 B/op	17 allocs/op
gorm_prep:	2857	414621 ns/op	5104 B/op	56 allocs/op
ent:		2830	432919 ns/op	4725 B/op	98 allocs/op
pgx_pool:	2952	436289 ns/op	305 B/op	8 allocs/op
beego:		2688	437135 ns/op	1736 B/op	46 allocs/op
pop:		2608	456999 ns/op	5744 B/op	170 allocs/op
dbr:		2607	461148 ns/op	2651 B/op	57 allocs/op
jet:		3121	461299 ns/op	4511 B/op	119 allocs/op
bun:		2838	462722 ns/op	4761 B/op	5 allocs/op
reform:		2833	463881 ns/op	1776 B/op	51 allocs/op
gorm:		2593	506949 ns/op	6864 B/op	99 allocs/op
gen:		2238	508708 ns/op	13856 B/op	163 allocs/op
zorm:		2372	525872 ns/op	4384 B/op	76 allocs/op
rel:		1986	587468 ns/op	3048 B/op	45 allocs/op
godb:		1899	602621 ns/op	4953 B/op	130 allocs/op
xorm:		1734	634878 ns/op	4305 B/op	145 allocs/op
upper:		973	1261231 ns/op	16755 B/op	390 allocs/op

Read
pgx:		8109	154095 ns/op	776 B/op	18 allocs/op
sqlc:		7712	155252 ns/op	904 B/op	19 allocs/op
pgx_pool:	7197	156529 ns/op	963 B/op	19 allocs/op
raw:		7870	160305 ns/op	2093 B/op	50 allocs/op
beego:		7533	161155 ns/op	2112 B/op	75 allocs/op
reform:		7388	166784 ns/op	3230 B/op	86 allocs/op
gorp:		7672	171981 ns/op	3332 B/op	122 allocs/op
pop:		6888	172396 ns/op	3189 B/op	66 allocs/op
ent:		6969	178647 ns/op	5685 B/op	145 allocs/op
sq:		6478	179199 ns/op	11080 B/op	126 allocs/op
pg:		7084	182878 ns/op	872 B/op	20 allocs/op
dbr:		7016	183563 ns/op	2184 B/op	36 allocs/op
bun:		6620	184279 ns/op	5844 B/op	39 allocs/op
sqlboiler:	7078	184351 ns/op	951 B/op	14 allocs/op
rel:		6942	188236 ns/op	2336 B/op	47 allocs/op
gorm_prep:	6218	189156 ns/op	4598 B/op	89 allocs/op
zorm:		6360	197832 ns/op	3265 B/op	65 allocs/op
jet:		6078	200318 ns/op	12858 B/op	249 allocs/op
gorm:		5424	227567 ns/op	5013 B/op	102 allocs/op
gen:		4958	249278 ns/op	10955 B/op	156 allocs/op
sqlx:		3729	331710 ns/op	2008 B/op	43 allocs/op
godb:		3574	348452 ns/op	4033 B/op	94 allocs/op
upper:		3331	358062 ns/op	5087 B/op	110 allocs/op
xorm:		3320	365678 ns/op	5161 B/op	131 allocs/op

ReadSlice
reform:		7723	164620 ns/op	4044 B/op	100 allocs/op
pgx_pool:	4611	247969 ns/op	30380 B/op	513 allocs/op
pgx:		4549	248481 ns/op	30320 B/op	513 allocs/op
sqlc:		4585	259026 ns/op	54625 B/op	620 allocs/op
raw:		4021	288801 ns/op	38374 B/op	1038 allocs/op
pg:		3510	341693 ns/op	22988 B/op	629 allocs/op
gorp:		3253	353983 ns/op	56968 B/op	1422 allocs/op
sqlx:		3152	365126 ns/op	37512 B/op	1225 allocs/op
upper:		3348	365565 ns/op	4822 B/op	90 allocs/op
ent:		3238	369594 ns/op	77398 B/op	2036 allocs/op
pop:		3208	370882 ns/op	76636 B/op	1306 allocs/op
sq:		2941	391981 ns/op	152430 B/op	1820 allocs/op
dbr:		2823	412680 ns/op	30881 B/op	1245 allocs/op
beego:		2779	414816 ns/op	55350 B/op	3077 allocs/op
sqlboiler:	2781	420416 ns/op	58932 B/op	1260 allocs/op
bun:		2792	423544 ns/op	34213 B/op	1124 allocs/op
gorm_prep:	2582	451883 ns/op	43596 B/op	2082 allocs/op
gorm:		2378	501733 ns/op	44804 B/op	2196 allocs/op
gen:		2004	538621 ns/op	50783 B/op	2250 allocs/op
zorm:		2109	553351 ns/op	163127 B/op	2158 allocs/op
jet:		2116	559535 ns/op	184339 B/op	2828 allocs/op
xorm:		2080	563435 ns/op	121235 B/op	4407 allocs/op
godb:		1899	632865 ns/op	68991 B/op	2284 allocs/op
rel:		1760	666736 ns/op	149044 B/op	2553 allocs/op
```

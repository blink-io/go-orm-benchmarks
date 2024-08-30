# Results

- orm-benchmark (with no flags)
```
Reports:

Insert
raw:		4015	301055 ns/op	704 B/op	13 allocs/op
pgx_pool:	3726	302171 ns/op	288 B/op	10 allocs/op
pgx:		4616	313061 ns/op	271 B/op	10 allocs/op
reform:		4032	318961 ns/op	1775 B/op	51 allocs/op
beego:		3936	320792 ns/op	2400 B/op	57 allocs/op
sqlc:		3670	321071 ns/op	2788 B/op	62 allocs/op
jet:		4120	325972 ns/op	3575 B/op	105 allocs/op
ent:		4050	327519 ns/op	4142 B/op	97 allocs/op
sqlboiler:	3589	328191 ns/op	1574 B/op	34 allocs/op
gorp:		4425	329747 ns/op	1782 B/op	41 allocs/op
sq:		4140	333445 ns/op	9761 B/op	100 allocs/op
dbr:		3591	344084 ns/op	2688 B/op	65 allocs/op
gorm_prep:	3804	344650 ns/op	5176 B/op	65 allocs/op
pg:		3538	353444 ns/op	1111 B/op	10 allocs/op
bun:		3364	359791 ns/op	5011 B/op	14 allocs/op
gorm:		3463	395021 ns/op	7176 B/op	105 allocs/op
gen:		3423	406706 ns/op	10057 B/op	134 allocs/op
sqlx:		2836	488875 ns/op	856 B/op	19 allocs/op
zorm:		2667	505117 ns/op	3799 B/op	77 allocs/op
upper:		2498	527069 ns/op	5895 B/op	125 allocs/op
xorm:		2479	533038 ns/op	3311 B/op	88 allocs/op
rel:		1831	547150 ns/op	2622 B/op	45 allocs/op
godb:		2497	552649 ns/op	4521 B/op	115 allocs/op
pop:		1630	749875 ns/op	9576 B/op	237 allocs/op

InsertMulti
pgx_pool:	1383	927008 ns/op	112945 B/op	42 allocs/op
raw:		1354	981957 ns/op	183805 B/op	930 allocs/op
pgx:		1304	1010004 ns/op	112902 B/op	42 allocs/op
beego:		1272	1030655 ns/op	177689 B/op	2745 allocs/op
sq:		985	1119976 ns/op	227280 B/op	1706 allocs/op
gorm_prep:	1098	1125536 ns/op	251037 B/op	1890 allocs/op
reform:		1119	1209480 ns/op	458844 B/op	2746 allocs/op
ent:		901	1261483 ns/op	386565 B/op	4598 allocs/op
jet:		888	1397326 ns/op	327550 B/op	6493 allocs/op
pg:		873	1475343 ns/op	3315 B/op	112 allocs/op
sqlx:		859	1476052 ns/op	170383 B/op	1551 allocs/op
gen:		756	1537763 ns/op	304645 B/op	5358 allocs/op
bun:		853	1567037 ns/op	42599 B/op	219 allocs/op
gorm:		728	1599345 ns/op	291412 B/op	5231 allocs/op
rel:		727	1726475 ns/op	306914 B/op	3265 allocs/op
zorm:		748	1783623 ns/op	199928 B/op	2780 allocs/op
upper:		676	1785427 ns/op	322889 B/op	4204 allocs/op
godb:		687	1832797 ns/op	254066 B/op	5894 allocs/op
xorm:		734	1899932 ns/op	248188 B/op	5414 allocs/op
dbr:		bulk-insert is not supported
sqlc:		bulk-insert is not supported
pop:		bulk-insert is not supported
sqlboiler:	bulk-insert is not supported
gorp:		bulk-insert is not supported

Update
raw:		8350	150441 ns/op	750 B/op	13 allocs/op
sqlc:		8707	152750 ns/op	878 B/op	14 allocs/op
sqlx:		3890	318564 ns/op	872 B/op	20 allocs/op
pgx:		4353	320560 ns/op	268 B/op	10 allocs/op
reform:		4195	322520 ns/op	1774 B/op	51 allocs/op
sqlboiler:	4216	323107 ns/op	901 B/op	17 allocs/op
ent:		3829	325074 ns/op	4677 B/op	97 allocs/op
sq:		3871	328377 ns/op	7415 B/op	90 allocs/op
beego:		3738	334986 ns/op	1752 B/op	47 allocs/op
pgx_pool:	3798	345384 ns/op	286 B/op	10 allocs/op
gorm_prep:	4087	347788 ns/op	5008 B/op	56 allocs/op
pop:		3475	349964 ns/op	6049 B/op	186 allocs/op
jet:		4090	366399 ns/op	4558 B/op	126 allocs/op
gorp:		4074	368648 ns/op	1206 B/op	32 allocs/op
dbr:		3448	369701 ns/op	2651 B/op	57 allocs/op
pg:		3456	373451 ns/op	768 B/op	9 allocs/op
bun:		3782	382205 ns/op	4729 B/op	5 allocs/op
gorm:		3405	394137 ns/op	6752 B/op	99 allocs/op
gen:		3175	424438 ns/op	13392 B/op	165 allocs/op
zorm:		2538	511633 ns/op	3024 B/op	59 allocs/op
xorm:		2062	515705 ns/op	3945 B/op	132 allocs/op
rel:		2715	533336 ns/op	3048 B/op	45 allocs/op
godb:		2506	545613 ns/op	5129 B/op	154 allocs/op
upper:		987	1217890 ns/op	16699 B/op	390 allocs/op

Read
pgx:		8368	148873 ns/op	892 B/op	8 allocs/op
pgx_pool:	7893	150838 ns/op	1078 B/op	9 allocs/op
raw:		8584	156398 ns/op	2077 B/op	50 allocs/op
sqlc:		8386	158899 ns/op	2092 B/op	51 allocs/op
beego:		7633	160386 ns/op	2112 B/op	76 allocs/op
reform:		7935	165779 ns/op	3215 B/op	86 allocs/op
gorp:		8107	168089 ns/op	3893 B/op	194 allocs/op
pop:		7822	169295 ns/op	3216 B/op	67 allocs/op
ent:		7052	175527 ns/op	5620 B/op	144 allocs/op
pg:		6954	176013 ns/op	872 B/op	20 allocs/op
sq:		7552	176762 ns/op	11143 B/op	135 allocs/op
bun:		7159	177804 ns/op	5829 B/op	39 allocs/op
dbr:		7480	180729 ns/op	2168 B/op	36 allocs/op
gorm_prep:	7076	180852 ns/op	4422 B/op	87 allocs/op
sqlboiler:	7267	181814 ns/op	951 B/op	14 allocs/op
rel:		6657	182035 ns/op	2320 B/op	47 allocs/op
zorm:		6805	192149 ns/op	3032 B/op	64 allocs/op
jet:		6192	199112 ns/op	12939 B/op	273 allocs/op
gorm:		5523	225234 ns/op	4791 B/op	98 allocs/op
gen:		4999	245695 ns/op	10334 B/op	153 allocs/op
sqlx:		4060	326776 ns/op	1992 B/op	43 allocs/op
godb:		3879	338710 ns/op	4081 B/op	102 allocs/op
upper:		3704	346996 ns/op	5069 B/op	110 allocs/op
xorm:		3762	353573 ns/op	4665 B/op	127 allocs/op

ReadSlice
reform:		7836	163406 ns/op	4029 B/op	100 allocs/op
pgx:		4773	249608 ns/op	42949 B/op	504 allocs/op
pgx_pool:	4621	253787 ns/op	43006 B/op	504 allocs/op
raw:		3793	308163 ns/op	38357 B/op	1038 allocs/op
sqlc:		3486	333367 ns/op	62676 B/op	1150 allocs/op
pg:		3415	341165 ns/op	22569 B/op	629 allocs/op
upper:		3342	364100 ns/op	4807 B/op	90 allocs/op
gorp:		3194	372993 ns/op	57401 B/op	1494 allocs/op
sqlx:		3142	373575 ns/op	37496 B/op	1225 allocs/op
ent:		3034	381962 ns/op	77204 B/op	2035 allocs/op
pop:		2950	387464 ns/op	68999 B/op	1306 allocs/op
sq:		2809	414263 ns/op	144935 B/op	1829 allocs/op
dbr:		2845	417456 ns/op	30800 B/op	1253 allocs/op
bun:		2725	430001 ns/op	34062 B/op	1124 allocs/op
beego:		2524	440292 ns/op	55218 B/op	3078 allocs/op
sqlboiler:	2680	442238 ns/op	66715 B/op	2260 allocs/op
gorm_prep:	2499	472148 ns/op	43193 B/op	2081 allocs/op
gorm:		2232	525686 ns/op	44349 B/op	2191 allocs/op
gen:		2011	552854 ns/op	49926 B/op	2246 allocs/op
zorm:		2060	589188 ns/op	161633 B/op	2949 allocs/op
jet:		1950	606021 ns/op	199870 B/op	3642 allocs/op
rel:		1827	629687 ns/op	100656 B/op	2253 allocs/op
godb:		1852	648639 ns/op	75248 B/op	3084 allocs/op
xorm:		1808	658882 ns/op	119402 B/op	4401 allocs/op
```

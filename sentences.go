package main

type sentences struct {
	test           string
	fornam         string
	musikPåDistans string
}

func getSentences() sentences {
	fornam := `_Berättelser om förnamn. __Enligt namnlagen ska alla ha ett förnamn
	, och förnamnen får vara högst fyra till antalet. Många har traditionsenligt fått sitt namn, för att hedra minnet av släktingar som har gått bort. Det är också vanligt att namnen följer tidsandan. Ändå tror många föräldrar att de väljer unika namn för sina barn. Senast när barnet börjar på dagis märker de plötsligt att namnet är allmänt populärt, __och där finns fyra Emmor; tre Linneor; fem William, __och alla i personalen heter Monica.

	Vad betyder ett namn för den som bär det, för den som ger det och för den som ser eller hör det? __Fyra personer berättar om förnamn:

	_Mamma till Nova:
	_Min dotter heter Nova. __Jag ville ha ett namn som inte är så vanligt i Finland, __och som kan associeras med olika saker. Själv tänkte jag på en supernova, __något sorts mirakel. Jag har nämligen en sjukdom som påverkar min aorta, och därför var hela graviditeten ett stort frågetecken, __och inte alls någon självklarhet.

	_Filippa:
	__Jag fick förnamnet Filippa efter att min mamma på radion hört en flicka, __med namnet Filippa ringa in och önska en låt. Då tyckte min mamma att om barnet hon väntar är en flicka ska hon få heta Filippa. Det roliga är ju sen att jag dessutom blev musikälskare.

	_Gunnar:
	_Jag heter Gunnar. __Jag tyckte; då jag var liten, __att det bara var gamla gubbar som hette så. Jag fick mitt namn på grund av de gamla vikingarna och deras sagor. __Min mamma blev förtjust i namnet. __Själv fick jag ”tåla och lida”. Nå; det var ju bara jag som tyckte namnet var för ovanligt, __för jag hade velat smälta in. Jag tänkte inte på att också de andra barnen hade så att säga egna namn.
	Senare i livet har jag skämtat kring namnet, till exempel att uttrycket ”kunnari” i Finlands nationalsport boboll kommer direkt ur historien. __Då vikingarna anlände till Finlands kuster, och skeppen befolkades av en hel del personer som ofta hette just Gunnar, så försvarade sig lokalbefolkningen med olika redskap. Det kunde vara vad som helst som de hittade, __till exempel grenar och stenar. __Därifrån kommer uttrycket ”lyödä kunnari”.

	_Mamma till Milja:
	Jag bestämde mig för länge sedan att om jag får ett barn, ska hen kunna säga namnet i telefon till någon utan att behöva göra en massa påpekanden. Jag måste till exempel alltid säga "Jag heter Erika med ett e och k", och sedan har jag ett efternamn som inte är det lättaste heller. Jag ville undvika att mitt barn ska råka ut för samma sak. Så när jag fick min dotter döpte vi henne till det finaste namnet jag vet. _Hon heter Milja.`

	musikPåDistans := `Vasa övningsskolas gymnasiekör har hoppat på trenden med körsång på distans nu i coronatider. Gymnasiekören har spelat in en egen version av Aftonens sång.

	Medlemmarna har spelat in sina egna stämmor hemma och allt har mixats ihop till en musikvideo som nu finns på Youtube.
	
	Johan Sundqvist, timlärare i musik vid Vasa övningsskola, berättar att gymnasiets tidigare rektor Ulla Granfors tog kontakt och frågade om det fanns någon inspelning av Aftonens sång som kunde spelas på hennes pappa Thor Österholms begravning. Thor Österholm, som avled i mars, var tidigare rektor för Vasa övningsskola.
	
	Vi hade några inspelningar från skolfester men de hade dålig inspelningskvalitet, _säger Sundqvist.
	
	Inspirerad av andra skolkörer som under undantagstillståndet har spelat in verk på distans kom Sundqvist fram till att man kunde pröva samma koncept med Vasa övningsskolas gymnasiekör.

	Aftonens sång är en lugn och stillsam sång som egentligen är idealisk som virtuellt projekt.
	
	_Spelades in med mobiltelefoner.
	När en hel kör ska spela in en video på distans görs det inte i realtid eftersom det är svårt att tajma allas sång på grund av varierande nätuppkopplingar. Därför spelade de 36 medlemmarna i gymnasiekören in sina egna stämmor hemma med sina mobiltelefoner.
	
	Till sin hjälp hemma hade studerandena dels en metronom som Sundqvist spelat in för att ange tempot, och dels de olika stämmorna som separata ljudfiler. Sedan mixades allas bidrag ihop till en helhet.
	
	Sundqvist hade bett alla i kören klä sig i svart eftersom videon skulle spelas på Thor Österholms begravning. Efter begravningen lades videon också ut på Youtube där den fått positiv respons.
	
	Både Sundqvist och medlemmarna i gymnasiekören är mer än nöjda med slutresultatet.
	
	Jag trodde inte att man med telefonbandningar kunde åstadkomma något sådant här. Det var en intensiv process på ett par veckor där vi lärde oss mycket, både jag och studerandena.
	`
	return sentences{test: "Geocaching är modern skattsökning som utövas internationellt. Med hjälp av gps-mottagare eller karta söker man efter skatter som geocachare har gömt i en ask på intressanta ställen.", fornam: fornam, musikPåDistans: musikPåDistans}

}

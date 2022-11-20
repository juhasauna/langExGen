package main

type sentences struct {
	test                   string
	fornam                 string
	musikPåDistans         string
	enDagIAmbulansen       string
	undantagstillstånd     string
	småferetagIEnLitenStad string
}

func getSentences() sentences {
	// skämtat kring - joke about
	// redskap - tools
	// döpte - babtized
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

	// bidrag - contribution
	// Undantagstillstånd - state of emergency
	// nätuppkopplingar - network connections
	// stämmor - parts
	// hade bett - had asked
	// lades  posted
	// telefonbandningar - phone tapes
	// åstadkomma - achieve
	musikPåDistans := `Vasa övningsskolas gymnasiekör har hoppat på trenden med körsång på distans nu i coronatider. Gymnasiekören har spelat in en egen version av Aftonens sång.

	Medlemmarna har spelat in sina egna stämmor hemma och allt har mixats ihop till en musikvideo som nu finns på Youtube.
	
	__Johan Sundqvist; timlärare i musik vid Vasa övningsskola, berättar att gymnasiets tidigare rektor Ulla Granfors tog kontakt, och frågade om det fanns någon inspelning av Aftonens sång, som kunde spelas på hennes pappa Thor Österholms begravning. __Thor Österholm; som avled i mars, __var tidigare rektor för Vasa övningsskola.
	
	Vi hade några inspelningar från skolfester men de hade dålig inspelningskvalitet; säger Sundqvist.
	
	Inspirerad av andra skolkörer som under undantagstillståndet har spelat in verk på distans, kom Sundqvist fram till att man kunde pröva samma koncept med Vasa övningsskolas gymnasiekör.

	Aftonens sång är en lugn och stillsam sång som egentligen är idealisk som virtuellt projekt.
	
	_Spelades in med mobiltelefoner.
	När en hel kör ska spela in en video på distans görs det inte i realtid, eftersom det är svårt att tajma allas sång på grund av varierande nätuppkopplingar. Därför spelade de 36 medlemmarna i gymnasiekören in sina egna stämmor hemma med sina mobiltelefoner.
	
	Till sin hjälp hemma hade studerande-na dels en metronom som Sundqvist spelat in för att ange tempot, och dels de olika stämmorna som separata ljudfiler. Sedan mixades allas bidrag ihop till en helhet.
	
	Sundqvist hade bett alla i kören klä sig i svart eftersom videon skulle spelas på Thor Österholms begravning. Efter begravningen lades videon också ut på Youtube där den fått positiv respons.
	
	Både Sundqvist och medlemmarna i gymnasiekören är mer än nöjda med slutresultatet.
	
	Jag trodde inte att man med telefonbandningar kunde åstadkomma något sådant här. Det var en intensiv process på ett par veckor där vi lärde oss mycket, __både jag och studerande-na.
	`

	// vid något skede ändrade jag mig - at some stage changed my mind
	// oroliga anhöriga - worried relatives
	// hårfärgningen - hair dye
	// blodförgiftning - blood poisoning
	enDagIAmbulansen := `
	_Erblind Mustafa är ambulanssköterska. Direkt efter examen som sjuksköterska fick han jobb på akuten i Helsingborg, men visste redan då att han ville läsa vidare till ambulanssköterska. 
	
	Ambulansjobbet med högt och hektiskt tempo har alltid lockat mig. Tidigare vill jag bli polis men vid något skede ändrade jag mig. Jag gillar att komma till jobbet och inte veta vilka typer av patienter jag kommer att möta, _säger Erblind.

	Vi som jobbar i ambulansen har ofta mycket att göra hela arbetspasset, men vi har bara en patient åt gången och därför blir det automatisk lugnare än på akutmottagningen. Framkörningstiden till följande patient gör också att det finns lite tid för vila. Man behöver det lugnet för att kunna jobba när det sedan gäller, _berättar han.
	
	Inom ambulanssjukvården är det viktigt att kunna ställa om snabbt, __ena stunden kan det vara lugnt, men så händer plötsligt något och då måste man agera. Inte minst måste man kunna hantera oroliga anhöriga. __Personlighet spelar stor roll i yrket.
	
	Dagens första larm går och Erblind Mustafa hoppar in i ambulansen. En 85-årig kvinna; Dagmar; har värk i ryggen. Efter några undersökningar vill Erblind även testa hjärtats aktivitet. _Pulsen är normal. Ryggvärken kan ändå vara ett tecken på någon osynlig sjukdom, __och därför rekommenderar jag vidare undersökningar senare. Den här gången behöver du inte följa med in till sjukhuset, _konstaterar han. 
	
	Fint; Jag har så många hushållssysslor att göra. Och då kan jag också behålla min tid för hårfärgningen på fredag, _säger Dagmar lättat.
	
	På stationen blir det lite tid över mellan larmen. Kollegerna samlas i köket och har roligt tillsammans. Dagens sista larm gäller en man som har tagit för många tabletter. Medan mannen magpumpas på akutmottagningen i akutrum 1, kommer det in en äldre kvinna med blodförgiftning i akutrum 2. Med bara en skärm mellan sig kämpar båda för sina liv.
	`

	undantagstillstånd := `Undantagstillstånd är temporära rättsliga tillstånd när statsmakten tillfälligtvis ger sig själv utökade befogenheter för att ta kontroll över svåra eller farliga inhemska nödlägen, i synnerhet när dessa befogenheter inskränker människors friheter. Under ett undantagstillstånd sätts med andra ord de ordinära konstitutionella förfarandena ur spel, i syfte att återupprätta ordningen vid fara för rikets säkerhet, naturkatastrofer, eller andra nationella olyckstillbud. Om en stat formellt inför undantagstillstånd, vilket inte alla länder har konstitutionell möjlighet till, på nationell eller lokal nivå är det fråga om ett mycket akut nödläge.

	Termen "undantagstillstånd" kan emellertid också avse mindre allvarliga nödlägen som föranleder statsmakter att tillfälligt införa restriktioner mot människors friheter och rättigheter. Undantagstillstånd är ett etablerat rättsbegrepp för särskilda förhållanden under nationella nödlägen. Olika länders lagstiftningar kan emellertid också tala om nödtillstånd, kris, extraordinär händelse, katastroftillstånd, larmberedskap, belägringstillstånd, krigsrätt och krigstillstånd för delvis samma lägen.
	`
	return sentences{test: "Geocaching är modern skattsökning som utövas internationellt. Med hjälp av gps-mottagare eller karta söker man efter skatter som geocachare har gömt i en ask på intressanta ställen.", fornam: fornam, musikPåDistans: musikPåDistans, enDagIAmbulansen: enDagIAmbulansen, undantagstillstånd: undantagstillstånd, småferetagIEnLitenStad: småferetagIEnLitenStad}

}

const småferetagIEnLitenStad = `
__Läs reklamtexterna för fyra småföretag i en liten stad. __Välj det företag som bäst passar ihop med varje text. __Två alternativ är överflödiga.
_Vilket företag beskriver texten bäst? _Klädbutik; Inredningsbutik; Idrottscentrum; Skönhetssalong; Fastighetsmäklare; Campingplats.

Vi ligger nära havet; naturen och den populära Skärgårdens ringväg. Hos oss kan du övernatta i tält eller husvagn eller hyra en stuga. Stadens allmänna badstrand finns i närheten. Också stadscentrum med aktiviteter såsom äventyrsgolf, tennis; frisbeegolf och ridning finns på promenad och cykelavstånd. _Välkommen!

Som företagare vill jag hålla mig uppdaterad om de nyaste behandlingsformerna, _apparaterna och produkterna. Det är viktigt att kunderna betjänas med behandlingar och vård som är anpassade för deras individuella behov. På det sättet kan jag erbjuda mina kunder en avkopplande stund i vardagen. Vi säljer också hudvårdsprodukter av hög kvalitet. Boka tid och njut av lite lyx!

Vårt mål är att erbjuda produkter där kvalitet; design och komfort går hand i hand. Vi väljer bara kvalitetsmärken till vårt sortiment, och hos oss hittar du alternativ för både vardag och fest. _Urvalet är stort, också då det gäller storlekar och modeller som passar olika kroppsformer och kunder i olika åldrar. __Vill du ha personlig och vänlig service, _kom till oss!

_Är du trött på ditt hem? __Saknar du något nytt runt dig? Vi kan hjälpa dig både när du planerar en större renovering, och när du söker något mindre som kan pigga upp dig i vardagen. Vi tycker att föremål ska vara vackra men gärna ha en funktion. __Vi tycker också om att blanda olika stilar. Kom och bolla dina idéer med oss så hittar vi en bra lösning för dig!
`

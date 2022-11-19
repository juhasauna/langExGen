$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeFullPath = $cylanceWhiteListedFolder + "langExGen.exe"
if (Test-Path $exeFullPath) {
    rm $exeFullPath
}
go build -o  $exeFullPath; 
& $exeFullPath
    
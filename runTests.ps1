$wd = $pwd.ToString()
$packageName = "." + ($args[0] | Split-Path -Parent).ToLower().Replace($wd.ToLower(), "") + "\"

$testName = ($args[0] | Split-Path -Leaf) -replace "\.go$" -replace "_test$"
$testName = "Test_" + $testName + "$"

$cylanceWhiteListedFolder = "C:\MO\Motors_Generators\tech\SDC\"
$exeFullPath = $cylanceWhiteListedFolder + "langExGen.exe"

go test -v -run $testName $packageName -o $exeFullPath
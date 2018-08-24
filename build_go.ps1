# raspberry pi

Write-Output "Building for raspberry pi"

if (Test-Path ./build/raspberrypi) {
    rm -r -fo ./build/raspberrypi
}

$env:GOARM="7"
$env:GOOS="linux"
$env:GOARCH="arm"

if (Test-Path ./casa-server) {
    rm -r -fo ./casa-server
}
go build ./cmd/casa-server

if (Test-Path ./build/raspberrypi) {
    rm -r -fo ./build/raspberrypi
}
if (-Not (Test-Path ./build/raspberrypi)) {
    mkdir ./build/raspberrypi
}
Move-Item ./casa-server ./build/raspberrypi -Force

# windows

Write-Output "Building for windows"
if (Test-Path ./casa-server.exe) {
    rm -r -fo ./casa-server.exe
}

$env:GOOS="windows"
$env:GOARCH="386"

go build ./cmd/casa-server

if ((Test-Path ./build/windows)) {
    rm -r -fo ./build/windows
}
if (-Not (Test-Path ./build/windows)) {
    mkdir ./build/windows
}
Move-Item ./casa-server.exe ./build/windows/casa-server.exe -Force

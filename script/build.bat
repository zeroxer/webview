@echo off
setlocal

echo Prepare directories...
set script_dir=%~dp0
set script_dir=%script_dir:~0,-1%
set src_dir=%script_dir%\..
set build_dir=%script_dir%\..\build
mkdir "%build_dir%"

echo Webview directory: %src_dir%
echo Build directory: %build_dir%

:: If you update the nuget package, change its version here
set nuget_version=1.0.1418.22
echo Using Nuget Package microsoft.web.webview2.%nuget_version%
if not exist "%script_dir%\microsoft.web.webview2.%nuget_version%" (
	curl -sSLO https://dist.nuget.org/win-x86-commandline/latest/nuget.exe
	nuget.exe install Microsoft.Web.Webview2 -Version %nuget_version% -OutputDirectory %script_dir%
	echo Nuget package installed
)

echo Setting up environment for Go...
rem Argument quoting works for Go 1.18 and later but as of 2022-06-26 GitHub Actions has Go 1.17.11.
rem See https://go-review.googlesource.com/c/go/+/334732/
rem TODO: Use proper quoting when GHA has Go 1.18 or later.
set "CGO_CXXFLAGS=-I%script_dir%\microsoft.web.webview2.%nuget_version%\build\native\include"
set "CGO_LDFLAGS=-L%script_dir%\microsoft.web.webview2.%nuget_version%\build\native\x64"
set CGO_ENABLED=1

echo Building Go examples
mkdir build\examples\go
go build -o build\examples\go\issue857.exe examples\issue857.go || exit /b

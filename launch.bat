@echo off

:WAIT_LOOP
REM Check if MySQL is running
netstat -an | findstr ":3306.*LISTENING" >nul
IF ERRORLEVEL 1 (
    REM MySQL is not yet running, wait for 1 second and check again
    timeout /t 1 /nobreak >nul
    GOTO WAIT_LOOP
)


REM MySQL is running, wait for a moment before starting the Go application
timeout /t 2 /nobreak >nul

REM MySQL is running, start your Go application
START /B "" "C:\Users\faals\OneDrive\Escritorio\hotel-trivoli\hotel-trivoli-api\hotel-trivoli-api.exe"

:WAIT_GO_APP
REM Check if Go application is running on port 3001
netstat -an | findstr ":3001.*LISTENING" >nul
IF ERRORLEVEL 1 (
    REM Go application is not yet running, wait for 1 second and check again
    timeout /t 1 /nobreak >nul
    GOTO WAIT_GO_APP
)

echo "starting client app"
START "" "C:\Users\faals\AppData\Local\Hotel Trivoli\Hotel Trivoli.exe"
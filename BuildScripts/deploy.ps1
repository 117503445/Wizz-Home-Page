& "C:\User\App\docker-machine-Windows-x86_64\dm.exe" env ali | Invoke-Expression
docker rm WizzHomePage -f
docker build -t wizz-home-page ..
docker run --name WizzHomePage -d -p 8001:80 -v /var/www/data:/go/release/Wizz-Home-Page/data wizz-home-page
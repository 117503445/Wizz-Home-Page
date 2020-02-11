docker rm WizzHomePage -f
docker build -t wizz-home-page ..
docker run --name WizzHomePage -d -p 443:443 -v /var/www/data:/go/release/Wizz-Home-Page/data wizz-home-page
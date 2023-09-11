
export LC_ALL=en_US.UTF-8


# Enable the "universe" repository.
add-apt-repository --yes universe

sudo apt update

timedatectl set-timezone Africa/Lagos
apt --yes install locales-all


useradd --create-home --shell "/bin/bash" --groups sudo deployer
passwd --delete deployer
chage --lastday 0 deployer

rsync --archive --chown=deployer:deployer /root/.ssh /home/deployer

# Configure the firewall to allow SSH, HTTP and HTTPS traffic.
ufw allow 22
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable

apt --yes install fail2ban


# Install the migrate CLI tool.
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 /usr/local/bin/migrate


# Install PostgreSQL.
apt --yes install postgresql
sudo -i -u postgres psql -c "CREATE DATABASE deployer"
sudo -i -u postgres psql -d deployer -c "CREATE ROLE deployer_user_new WITH LOGIN PASSWORD 'xx992sKJS8291113'"
sudo -i -u postgres psql -d deployer -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp"'

# Install Caddy (see https://caddyserver.com/docs/install#debian-ubuntu-raspbian).
apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
apt update
apt --yes install caddy

# Install Redis.
apt --yes install redis-server
# Configure Redis to start on boot.
systemctl enable redis-server.service

# you can skip this if you want
reboot



# You may need to give the user a sudo permission
#sudo visudo
# deployer ALL=(ALL) NOPASSWD:ALL

# this is safer. Give access to fewer programs
# deployer ALL=(ALL) NOPASSWD:/usr/bin/systemctl,/usr/local/bin/migrate,/home/deployer/api

# Want to view the logs
# sudo journalctl -u api --follow

# 🥧 Raspberry Pi DevOps Lab

라즈베리파이 환경에서 구축한 3-Tier(Web-App-DB) 아키텍처 실습 프로젝트입니다.
Nginx(Reverse Proxy), Go(Backend), PostgreSQL(Database)을 Docker Compose로 구성하였습니다.

## 🏗 Architecture
- **Web Server**: Nginx (SSL/TLS 적용, Reverse Proxy)
- **App Server**: Go (Golang 1.2x, Hot-Reload with Air)
- **Database**: PostgreSQL 13
- **Infrastructure**: Docker & Docker Compose

## 🚀 Prerequisites (준비사항)
이 프로젝트를 실행하기 위해 다음 도구들이 필요합니다.
- Docker & Docker Compose
- SSL 인증서 생성 도구 (OpenSSL)

## 🛠 Installation & Setup

### 1. Repository Clone
```Bash
git clone [https://github.com/YOUR_ID/raspberry-devops-lab.git](https://github.com/YOUR_ID/raspberry-devops-lab.git)
cd raspberry-devops-lab
```

### 2. Environment Variables (.env) 설정
보안상 .env 파일은 포함되어 있지 않습니다. 프로젝트 루트에 직접 생성해야 합니다.
```Bash
# .env 파일 생성
echo "POSTGRES_USER=myuser" >> .env
echo "POSTGRES_PASSWORD=mypassword" >> .env
echo "POSTGRES_DB=mydb" >> .env
```

### 3. SSL 인증서 생성 (Self-Signed)
HTTPS 적용을 위한 사설 인증서를 생성합니다.
```Bash
mkdir certs
openssl req -new -newkey rsa:2048 -nodes \
  -keyout certs/my-priv.key -out certs/my-cert.crt \
  -x509 -days 365 -subj "/CN=my-devops.com"
```
### 4. Hosts 설정 (Local Only)
PC에서 접속하기 위해 hosts 파일을 수정해야 합니다.
- Windows: C:\Windows\System32\drivers\etc\hosts
- Mac/Linux: /etc/hosts

```Plaintext
# 라즈베리파이 IP로 변경하세요
192.168.0.xxx  my-devops.com
```

## ▶️ Usage (실행 방법)
```Bash
# 컨테이너 빌드 및 실행 (백그라운드)
docker-compose up -d --build

# 로그 확인 (Go 앱)
docker-compose logs -f go-app
```

## 🔗 Endpoints
- Main Page: https://my-devops.com
- API Test: https://www.google.com/search?q=https://my-devops.com/api
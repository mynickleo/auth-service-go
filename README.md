# ✨ Auth Service Backend

Это небольшой бэкенд-сервис для управления аутентификацией и авторизацией пользователей, написанный на Go. Сервис поддерживает регистрацию, вход в систему, управление пользователями и ролями. Этот проект разработан с использованием PostgreSQL для хранения данных, Redis для кеширования, Gomail для отправки электронных писем, и MinIO (S3-совместимое хранилище) для хранения изображений аватаров

> This is a small backend service for user authentication and authorization, written in Go. The service supports registration, login, user and role management. This project uses PostgreSQL for data storage, Redis for caching, Gomail for email sending, and MinIO (an S3-compatible storage) for storing user avatars

## 📦 Используемые технологии | Technologies Used
- Go
- PostgreSQL
- Redis
- Gomail
- Sqlc
- MinIO (S3): Bonus

## 🚀 Основные возможности | Key Features
- Регистрация и аутентификация пользователей
- Отправка писем для подтверждения и восстановления пароля
- Загрузка и хранение аватаров пользователей
- Кеширование данных с помощью Redis

> - User registration and authentication
> - Email notifications for verification and password reset
> - User avatar upload and storage
> - Data caching with Redis

## 📂 Структура проекта | Project Structure
- `cmd/server/`
- `config/`
- `internal/`
- - `app/`
- - `controllers/`
- - `database/`
- - - `postgres/`
- - - - `migrations/`
- - - - `queries/`
- - - `redis/`
- - `interfaces/`
- - `models/`
- - `module/`
- - `repository/`
- - `services/`
- - `utils/`
- `pkg/`
- - `minio/`
- - `sqlcqueries/`

## ⚙️ Установка и запуск | Installation and Launch

```bash
git clone https://github.com/mynickleo/auth-service-go.git
cd auth-service-go
```

```bash
docker-compose up -d
```

```bash
go run cmd/server/main.go
```

## 🔗 API

**`Аутентификация | Authentication`**
- POST /api/auth/send-mail 
- POST /api/auth/register 
- POST /api/auth/login 

**`Системный статус | System Status`**
- GET /api/ready

**`Управление ролями | Role Management`**
> Проверка на jwt токен | Checking for a jwt token
- POST /api/roles
- POST /api/user-roles
- GET /api/roles
- GET /api/roles/:id
- GET /api/user-role/:id
- PUT /api/roles/:id
- DELETE /api/roles

**`Управление пользователями | User Management`**
> Проверка на jwt токен | Checking for a jwt token
- POST /api/users
- POST /api/users/avatar
- GET /api/users
- GET /api/users/:id
- PUT /api/users/:id
- DELETE /api/users/:id

## 📬 Обратная связь | Feedback
Если у вас есть предложения или вы нашли ошибки, создайте Issue или отправьте Pull Request. Если вам понравилось, то можете дать ⭐ этому репозиторию
> If you have suggestions or find any issues, feel free to open an Issue or submit a Pull Request! If you liked it, you can give ⭐ to this repository
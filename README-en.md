# NoFap Helper (戒色助手)

<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width="300" height="300" />
</div>

<div align=center>
<img src="https://img.shields.io/badge/golang-1.20-blue"/>
<img src="https://img.shields.io/badge/gin-1.9.1-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.3.4-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.3.8-green"/>
<img src="https://img.shields.io/badge/gorm-1.25.2-red"/>
<img src="https://img.shields.io/badge/uni--app-3.0-orange"/>
</div>

<div align=center>
<a href="https://trendshift.io/repositories/3250" target="_blank"><img src="https://trendshift.io/api/badge/repositories/3250" alt="Calcium-Ion%2Fnew-api | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>
</div>

English | [简体中文](./README.md)

## 🎯 Project Overview

**NoFap Helper** is a gamified health management application focused on helping young people overcome pornography addiction. Through scientific assessment systems, gamified incentive mechanisms, and community support features, it provides comprehensive support for users' recovery journey.

### 🌟 Core Features

- **Scientific Assessment**: Based on internationally recognized sexual addiction assessment scales (SAST-R, PATHOS)
- **Gamified Experience**: 50-level system + 100+ achievement badges + experience point rewards
- **Community Support**: Anonymous community environment providing emotional support and practical advice
- **Emergency Help**: One-click emergency assistance for critical moments
- **Personalized Learning**: AI recommendation system providing customized recovery content

## 🏗️ Technical Architecture

### Frontend Technology Stack
- **Mini Program**: uni-app X + Vue 3 + TypeScript + Pinia
- **Admin Panel**: Vue 3 + Element Plus + Vite
- **UI Design**: Custom component library based on Tailwind CSS
- **State Management**: Pinia
- **Development Tools**: Vite + TypeScript

### Backend Technology Stack
- **Primary Language**: Go + Gin + GORM
- **Database**: MySQL 8.0+ + Redis
- **Architecture Pattern**: MVC architecture
- **Authentication**: JWT
- **API Documentation**: Auto-generated Swagger

## 🚀 Quick Start

### Environment Requirements
```
- Node.js 18+
- Go 1.22+
- MySQL 8.0+
- Redis 6.0+
- WeChat Developer Tools
```

### Frontend Development

#### Mini Program
```bash
cd frontend
npm install
npm run dev:mp-weixin  # WeChat Mini Program development
npm run dev:h5         # H5 development and debugging
```

#### Admin Panel
```bash
cd web
npm install
npm run serve
```

### Backend Development
```bash
cd server
go mod tidy
go run main.go
```

### Database Initialization
```bash
# Connect to MySQL database
mysql -u root -p123456

# Create database
CREATE DATABASE gva_nofap CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Run database migration scripts
cd server
go run main.go
```

## 📱 Core Features

### 1. Pornography Addiction Index Assessment System
- 50-question scientific assessment questionnaire
- 5-level risk assessment
- Assessment history tracking
- Regular re-assessment reminders

### 2. Gamified Incentive System
- **Level System**: 50 levels corresponding to different recovery milestones
- **Experience Points**: Daily check-ins, task completion, community contributions
- **Achievement System**: 100+ achievement badges
- **Virtual Rewards**: Theme skins, exclusive avatars, personalized titles

### 3. Daily Check-in Function
- Daily check-in records
- Emotional state assessment
- Consecutive days statistics
- Check-in history viewing

### 4. Community Support Features
- Anonymous posting system
- Content categorization management
- Like and comment functionality
- AI + manual content moderation

### 5. Emergency Help System
- One-click emergency assistance
- Attention diversion activities
- Community volunteer response
- Professional resource recommendations

### 6. Learning and Growth Module
- Personalized content recommendations
- Articles/videos/audio content
- Learning progress tracking
- Content bookmarking functionality

### 7. Data Analysis Reports
- Progress statistics visualization
- Emotional change trends
- Success rate analysis
- Personal growth reports

## 🎨 Design Standards

### Visual Design Principles
- **Warm and Friendly**: Use warm color combinations, avoid cold feelings
- **Private and Secure**: Interface design reflects privacy protection, reducing user psychological pressure
- **Gamified Visual**: Incorporate gaming elements while maintaining professionalism
- **Simple and Clear**: Clear information hierarchy, simple operation paths

### Color Standards
- **Primary Color**: #34D399 (Warm emerald green)
- **Secondary Color**: #06B6D4 (Fresh sky blue)
- **Accent Color**: #F59E0B (Vibrant amber)
- **Background Color**: #F8FAFC (Light gray-white, eye-friendly)
- **Emergency Color**: #EF4444 (Vibrant red, only for emergency buttons)

## 📊 Project Structure

```
gva_NoFap/
├── frontend/                 # Frontend mini program code
│   ├── src/
│   │   ├── pages/           # Page files
│   │   │   ├── welcome/     # Welcome page
│   │   │   ├── assessment/  # Assessment page
│   │   │   ├── checkin/     # Check-in page
│   │   │   ├── community/   # Community page
│   │   │   ├── emergency/   # Emergency help
│   │   │   ├── learning/    # Learning page
│   │   │   ├── progress/    # Progress page
│   │   │   └── profile/     # Personal center
│   │   ├── components/      # Components
│   │   ├── store/          # State management
│   │   ├── api/            # API interfaces
│   │   └── utils/          # Utility functions
│   ├── manifest.json       # App configuration
│   ├── pages.json         # Page routing configuration
│   └── package.json
├── server/                  # Backend service code
│   ├── api/               # API controllers
│   │   ├── v1/
│   │   │   ├── miniprogram/  # Mini program APIs
│   │   │   └── system/       # System management APIs
│   ├── model/             # Data models
│   │   ├── miniprogram/   # Mini program business models
│   │   └── system/        # System management models
│   ├── service/           # Business logic
│   ├── router/            # Routing configuration
│   ├── middleware/        # Middleware
│   └── config/            # Configuration files
├── web/                    # Admin panel code
│   ├── src/
│   │   ├── view/          # Pages
│   │   ├── components/    # Components
│   │   ├── api/          # API interfaces
│   │   └── router/       # Routing
└── docs/                  # Project documentation
    ├── PRD.md            # Product Requirements Document
    ├── Database_Schema.md # Database Design Document
    └── API.md            # API Interface Document
```

## 🗄️ Database Design

### Core Business Tables
1. **Users Table** (users) - User basic information
2. **Abstinence Records Table** (abstinence_records) - Recovery progress tracking
3. **Assessment Results Table** (assessment_results) - Addiction index assessment
4. **Daily Check-ins Table** (daily_checkins) - Check-in records
5. **Community Posts Table** (community_posts) - Community content
6. **Achievements Table** (achievements) - Achievement system
7. **Learning Contents Table** (learning_contents) - Learning resources
8. **Emergency Requests Table** (emergency_requests) - Help request records

## 📈 Product Metrics

### Core KPIs
- **Daily Active Users (DAU)**: Target 10,000+ (within 6 months)
- **User Retention Rate**: Next-day retention >70%, 7-day retention >45%
- **Recovery Success Rate**: 30-day success rate >60%, 90-day success rate >35%
- **Community Activity**: Daily posts >500, user participation rate >40%

### North Star Metric
**Total cumulative recovery days of all users** - Directly reflects the core value of helping users achieve their recovery goals

## 🔧 Development Standards

### Code Standards
- Frontend follows Vue 3 Composition API standards
- Backend follows Go standard code conventions
- Unified use of TypeScript/Go strict type checking

### Commit Standards
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation updates
- `style`: Code formatting adjustments
- `refactor`: Code refactoring
- `test`: Testing related
- `chore`: Build process or auxiliary tool changes

## 🚀 Deployment Guide

### Development Environment
- **Frontend**: HBuilderX + WeChat Developer Tools
- **Backend**: GoLand/VSCode + MySQL + Redis
- **Admin Panel**: VSCode + Chrome

### Production Environment
- **Server**: Tencent Cloud/Aliyun
- **Database**: Cloud Database MySQL
- **Cache**: Cloud Database Redis
- **CDN**: Tencent Cloud COS/Aliyun OSS
- **Deployment**: Docker + Docker Compose

## 📋 Development Tasks

This project uses TaskMaster AI for task management, including 25 main development tasks:

### Current Progress
- ✅ **Task 1: Project Foundation Architecture Setup** - Completed
- ✅ **Task 2: Database Design and Implementation** - Completed
- 🔄 **Task 3: User Authentication System** - In Progress

### Complete Task List
1. Project Foundation Architecture Setup ✅
2. Database Design and Implementation ✅
3. User Authentication System 🔄
4. Frontend Basic UI Component Library
5. Welcome Page Implementation
6. Addiction Index Assessment System
7. Daily Check-in Function
8. Gamified Incentive System
9. Home Dashboard Implementation
10. Progress Tracking Page
11. Community Support Features
12. Community Page UI Implementation
13. Emergency Help System
14. Emergency Help Page UI
15. Learning Content Management System
16. Learning Page UI Implementation
17. Personal Center Features
18. Personal Center Page UI
19. Admin Panel User Management Features
20. Admin Panel Content Management Features
21. API Documentation and Testing
22. Security Hardening and Optimization
23. Mini Program Review Preparation
24. Deployment and Launch
25. User Testing and Feedback Collection

## 🤝 Contributing

We welcome all forms of contributions! Please read our contributing guidelines:

1. Fork this repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 📞 Contact Us

- **Project Homepage**: [https://github.com/your-username/gva_NoFap](https://github.com/your-username/gva_NoFap)
- **Issue Reports**: [Issues](https://github.com/your-username/gva_NoFap/issues)
- **Feature Suggestions**: [Discussions](https://github.com/your-username/gva_NoFap/discussions)

## 🙏 Acknowledgments

Thank you to all developers and users who have contributed to this project!

---

**NoFap Helper** - Making recovery simple and fun 🎯


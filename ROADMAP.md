# Roadmap Pengembangan Aplikasi Pengelolaan Keuangan (LokaApp - MVP)

Dokumen ini berisi rencana pengembangan aplikasi keuangan digital yang terstruktur, aman, dan scalable.

## 🎯 Objective (Tujuan Utama)
1.  **Budgeting Berbasis Kategori**: Membagi pendapatan ke dalam kategori (Food, Rent, Savings, dll).
2.  **Multi-Source Tracking**: Memisahkan saldo dan transaksi berdasarkan sumber dana (Bank, E-wallet, RDN, dll).
3.  **Data Integrity**: Menjamin akurasi saldo dengan sistem Database Transaction (ACID).
4.  **Financial Health Radar**: Visualisasi sisa saldo dan batas aman pengeluaran.
5.  **Security First**: Proteksi data pengguna dengan standar enkripsi modern.

---

## 🗺️ Roadmap Pengembangan

### Fase 1: Fondasi, Backend & Database (Minggu 1-2)
Fokus pada integritas data, keamanan, dan logika bisnis.
- [ ] **Setup Lingkungan & Repository**
    - [x] Inisialisasi Git & Remote Repository
    - [x] Setup Docker Compose (Postgres & PGAdmin)
    - [x] Dokumentasi Awal (README.md & ROADMAP.md)
    - [x] Inisialisasi Project Go (Module setup & folder structure)
    - [ ] Konfigurasi Migration Tool (`golang-migrate`)
- [x] **Desain Database (Production Grade)**
    - [x] Schema `users`, `accounts`, `categories`, `budgets`, `transactions`
    - [x] Implementasi UUID & Enum Types
    - [x] Penambahan Indexing untuk performa
- [ ] **Logika Bisnis Core**
    - [ ] Implementasi DB Transaction (Atomic Updates)
    - [ ] Input Validation (Request DTOs)
    - [ ] Error Handling Middleware
- [ ] **API Core (with Scalability)**
    - [ ] Auth System (JWT + Refresh Token)
    - [ ] Accounts & Categories CRUD
    - [ ] Transactions CRUD (with Pagination & Filtering)
    - [ ] Dashboard Summary API

### Fase 2: Antarmuka Pengguna & Integrasi (Minggu 3-4)
Membangun UI yang cepat, cantik, dan responsif.
- [ ] **Frontend Setup**
    - [ ] Inisialisasi React (Vite) + Tailwind CSS
    - [ ] Integrasi shadcn/ui & Icons
    - [ ] Setup Axios & TanStack Query
- [ ] **Dashboard & Visualisasi**
    - [ ] Wallet/Account Selector Component
    - [ ] Budget Progress Tracking
    - [ ] Expense Charts (Recharts)
- [ ] **Feature Implementation**
    - [ ] Transaction Entry Form (Mobile-First)
    - [ ] Transaction History with Pagination
    - [ ] Category Management UI

### Fase 3: Quality Assurance & Cloud Deployment (Minggu 5)
- [ ] **Testing & Polish**
    - [ ] Unit Testing Backend (Logic calculation)
    - [ ] Integration Testing API
    - [ ] Bug fixing & Performance tuning
- [ ] **Deployment**
    - [ ] Setup GitHub Actions (CI/CD)
    - [ ] Deploy Backend (Railway/VPS)
    - [ ] Deploy Frontend (Vercel)

---

## 🗒️ Waiting List Fitur (Backlog)
- [ ] **Account Transfer**: Fitur pindah saldo antar rekening secara aman.
- [ ] **Telegram Bot**: Input cepat via chat.
- [ ] **Export Report**: Download PDF/CSV laporan bulanan.
- [ ] **Trading Journal**: Modul khusus SMC/Bandarmologi.

---
*Terakhir diperbarui: 4 Mei 2026*
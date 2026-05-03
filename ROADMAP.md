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
*   **Setup Lingkungan**:
    *   Inisialisasi Go dengan *Clean Architecture*.
    *   Setup Docker Compose (Postgres).
    *   Konfigurasi Migration Tool (`golang-migrate`).
*   **Desain Database (Production Grade)**:
    *   `users`: ID, email, hashed_password.
    *   `accounts`: ID, user_id, name, type, balance (Decimal/BigInt).
    *   `categories`: ID, user_id, name, icon, type.
    *   `budgets`: ID, category_id, amount, period.
    *   `transactions`: ID, user_id, category_id, account_id, amount, note, date.
*   **Logika Bisnis Core**:
    *   **Atomic Updates**: Implementasi DB Transaction untuk setiap pencatatan transaksi agar saldo rekening selalu akurat.
    *   **Input Validation**: Validasi ketat untuk mencegah input saldo negatif atau kategori ilegal.
*   **API Core (with Scalability)**:
    *   `ACCOUNTS & CATEGORIES`: CRUD lengkap.
    *   `TRANSACTIONS`: CRUD dengan fitur **Pagination** dan **Date-Range Filtering**.
    *   `AUTH`: JWT + Refresh Token logic.

### Fase 2: Antarmuka Pengguna & Integrasi (Minggu 3-4)
Membangun UI yang cepat, cantik, dan responsif.
*   **Stack**: React.js, Tailwind CSS, shadcn/ui.
*   **State Management**: TanStack Query (untuk caching & optimisctic updates).
*   **Mobile-First Design**: Memastikan input pengeluaran nyaman digunakan di layar HP.
*   **Dashboard Utama**:
    *   **Wallet Selector**: Menampilkan saldo per rekening.
    *   **Progress Bar Budget**: Visualisasi real-time per kategori.
*   **Visualisasi Data**: Chart distribusi pengeluaran mingguan/bulanan.

### Fase 3: Quality Assurance & Cloud Deployment (Minggu 5)
*   **Testing**: Unit testing untuk perhitungan saldo dan integrasi API.
*   **CI/CD**: GitHub Actions untuk otomatisasi testing & deployment.
*   **Deployment**: Railway/Render/VPS + Vercel.

---

## 🗒️ Waiting List Fitur (Backlog)
- [ ] **Account Transfer**: Fitur pindah saldo antar rekening secara aman.
- [ ] **Telegram Bot**: Input cepat via chat.
- [ ] **Export Report**: Download PDF/CSV laporan bulanan.
- [ ] **Trading Journal**: Modul khusus SMC/Bandarmologi.

---
*Roadmap ini adalah panduan hidup yang akan disesuaikan dengan kendala teknis di lapangan.*
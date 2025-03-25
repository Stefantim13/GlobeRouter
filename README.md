# 🌍 GlobeRouter

**GlobeRouter** este o aplicație de planificare a rutelor multimodale (avion, tren, autobuz etc.), ce oferă utilizatorilor posibilitatea de a găsi cele mai eficiente trasee în funcție de preț, durată sau număr de schimbări. Utilizatorii pot vizualiza traseele pe hartă, salva rute favorite și primi notificări pentru reduceri de preț.

---

## ✨ Descriere scurtă

GlobeRouter permite planificarea călătoriilor între două locații geografice folosind combinații de mijloace de transport. Aplicația calculează automat cele mai bune rute în funcție de:

- Prețul cel mai mic
- Durata cea mai scurtă
- Numărul minim de schimbări

---

## 📋 Backlog

### 1. Căutare și planificare rută

#### 🧩 User Story 1.1: Căutare rută între două locații
- [ ] Task: Implementare formular cu autocomplete pentru locații
- [ ] Task: Validare date introduse
- [ ] Task: Integrare API-uri de transport (ex: Skyscanner, Transport API)

#### 🧩 User Story 1.2: Suport pentru multiple mijloace de transport
- [ ] Task: Modelarea conexiunilor multimodale
- [ ] Task: Calcularea timpilor de transfer între mijloace de transport
- [ ] Task: Vizualizarea traseului complet

---

### 2. Filtrare și sortare rute

#### 🧩 User Story 2.1: Filtrare după preț
- [ ] Task: Calculare total cost rută
- [ ] Task: Interfață de filtrare și sortare

#### 🧩 User Story 2.2: Filtrare după durată
- [ ] Task: Calculare durată totală rută (inclusiv transferuri)
- [ ] Task: Interfață de filtrare și sortare

#### 🧩 User Story 2.3: Filtrare după numărul de schimbări
- [ ] Task: Algoritm pentru numărare legături
- [ ] Task: Setare filtru în UI

---

### 3. Interfață utilizator (UI/UX)

#### 🧩 User Story 3.1: Vizualizare rută pe hartă
- [ ] Task: Integrare hartă (ex: Leaflet, Mapbox, Google Maps)
- [ ] Task: Desenarea rutelor pe hartă

#### 🧩 User Story 3.2: Pagină cu detalii rută
- [ ] Task: Afișare tip transport, companie, durată, preț
- [ ] Task: Buton de rezervare (sau redirect către site-ul providerului)

---

### 4. Funcționalități extra _(opțional)_

#### 🧩 User Story 4.1: Salvare rute favorite
- [ ] Task: Salvare rută în baza de date
- [ ] Task: Pagină „Rutele mele”

#### 🧩 User Story 4.2: Notificări pentru schimbări de preț
- [ ] Task: Sistem de notificare pentru prețuri (email, push)

---

### 5. Autentificare și profil utilizator _(opțional)_

#### 🧩 User Story 5.1: Autentificare/înregistrare
- [ ] Task: Pagini de login/register

#### 🧩 User Story 5.2: Pagină de profil
- [ ] Task: Afișare rute salvate

---

## 🚀 Te interesează MVP-ul?

Recomandăm pentru început un MVP (Minimum Viable Product) cu:
- Căutare rută simplă (fără combinații)
- Filtru după preț
- Vizualizare pe hartă
- Interfață simplă pentru detalii rută

---

## 🛠️ Tehnologii recomandate

- **Frontend**: React + Leaflet / Mapbox
- **Backend**: Python (FastAPI) / Node.js
- **Bază de date**: PostgreSQL + PostGIS
- **Autentificare**: Firebase Auth / OAuth 2.0
- **API-uri**: Skyscanner, Flixbus, Transport API

---

## 📌 Status: În dezvoltare 🚧

Proiect: GlobeRouter – Aplicatie de planificare a rutelor multimodale (avion, tren, autobuz, etc.)
Descriere scurta: Aplicatia permite utilizatorilor sa planifice calatorii intre doua puncte geografice folosind diferite mijloace de transport (avion, tren, autobuz). Aceasta calculeaza rutele posibile in functie de criterii precum cel mai mic pret, cea mai scurta durata sau cele mai putine schimbari. Utilizatorii pot vizualiza traseele pe harta, salva rute favorite si primi notificari pentru reduceri de pret.

Backlog
1: Cautare si planificare ruta
User Story 1.1: Cautare ruta intre doua locatii
    • Task: Implementare formular cu autocomplete pentru locatii.
    • Task: Validare date introduse.
    • Task: Integrare API-uri de transport (ex: Skyscanner, Transport API).
User Story 1.2: Suport pentru multiple mijloace de transport
    • Task: Modelarea conexiunilor multimodale.
    • Task: Calcularea timpilor de transfer intre mijloace de transport.
    • Task: Vizualizarea traseului complet.

2: Filtrare si sortare rute
User Story 2.1: Filtrare dupa pret
    • Task: Calculare total cost ruta.
    • Task: Interfata de filtrare si sortare.
User Story 2.2: Filtrare dupa durata
    • Task: Calculare durata totala ruta (inclusiv transferuri).
    • Task: Interfata de filtrare si sortare.
User Story 2.3: Filtrare dupa numarul de schimbari
    • Task: Algoritm pentru numarare legaturi.
    • Task: Setare filtru in UI.

3: Interfata utilizator (UI/UX)
User Story 3.1: Vizualizare ruta pe harta
    • Task: Integrare harta (ex: Leaflet, Mapbox, Google Maps).
    • Task: Desenarea rutelor pe harta.
User Story 3.2: Pagina cu detalii ruta
    • Task: Afisare tip transport, companie, durata, pret.
    • Task: Buton de rezervare (sau redirect spre site-ul providerului).

4: Functionalitati extra - optional
User Story 4.1: Salvare rute favorite
    • Task: Salvare ruta in baza de date.
    • Task: Pagina "Rutele mele".
User Story 4.2: Notificari pentru schimbari de pret
    • Task: Sistem de notificare pentru preturi (email, push).

5: Autentificare si profil utilizator - optional
User Story 5.1: Autentificare/inregistrare
    • Task: Pagini de login/register.
User Story 5.2: Pagina de profil
    • Task: Afisare rute salvate.

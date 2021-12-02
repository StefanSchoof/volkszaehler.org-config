# volkszaehler.org-config

Dies ist meine [volkszaehler.org](https://www.volkszaehler.org/) config wie sie auf einem rasperby pi in meinem Haus läuft. Dies ist im Sinne von Linus Torvalds "Real men don’t use backups, they post their stuff on a public ftp server and let the rest of the world make copies."

Wenn es jemanden anderes dabei Hilfreich ist, ist dies natürlich auch toll.

# .env

Um traefik zu beniutzen muss eine `.env` Datei mit folgenden inhalten erstellt 
werden:

```
HOST=<traefik host>
```
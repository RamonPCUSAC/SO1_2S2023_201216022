CREATE DATABASE tarea3;
USE tarea3;
CREATE TABLE Albums (
    AlbumID INT AUTO_INCREMENT PRIMARY KEY,
    TituloAlbum VARCHAR(255),
    Artista VARCHAR(255),
    AnioLanzamiento VARCHAR(100),
    GeneroMusical VARCHAR(100)
);
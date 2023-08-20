import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [albums, setAlbums] = useState([]);
  const [albumData, setAlbumData] = useState({
    titulo_album: '',
    artista: '',
    anio_lanzamiento: '',
    genero_musical: '',
  });

  useEffect(() => {
    // Realiza una solicitud GET a la API al cargar la página
    fetch('http://localhost:8080/albums')
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setAlbums(data);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setAlbumData({ ...albumData, [name]: value });
  };

  const handleSubmit = (event) => {
    event.preventDefault();

    // Realiza una solicitud POST a la API para insertar un nuevo álbum
    fetch('http://localhost:8080/insert', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(albumData),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        // Recarga la lista de álbumes después de la inserción
        fetch('http://localhost:8080/albums')
          .then((response) => response.json())
          .then((data) => {
            console.log(data);
            setAlbums(data);
          })
          .catch((error) => {
            console.error(error);
          });
      })
      .catch((error) => {
        console.error(error);
      });

    // Limpia los campos del formulario
    setAlbumData({
      titulo_album: '',
      artista: '',
      anio_lanzamiento: '',
      genero_musical: '',
    });
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>SO1 - TAREA 3 - 2S2023</h1>
      </header>
      <div className="App-container">
        <div className="App-form">
          <h2>Registrar Álbum</h2>
          <form onSubmit={handleSubmit}>
            <div>
              <label>Título del Álbum:</label>
              <input
                type="text"
                name="titulo_album"
                value={albumData.titulo_album}
                onChange={handleInputChange}
              />
            </div>
            <div>
              <label>Artista:</label>
              <input
                type="text"
                name="artista"
                value={albumData.artista}
                onChange={handleInputChange}
              />
            </div>
            <div>
              <label>Año de Lanzamiento:</label>
              <input
                type="text"
                name="anio_lanzamiento"
                value={albumData.anio_lanzamiento}
                onChange={handleInputChange}
              />
            </div>
            <div>
              <label>Género Musical:</label>
              <input
                type="text"
                name="genero_musical"
                value={albumData.genero_musical}
                onChange={handleInputChange}
              />
            </div>
            <button type="submit">Guardar</button>
          </form>
        </div>
        <div className="App-table">
          <h2>Lista de Álbumes</h2>
          <table>
            <thead>
              <tr>
                <th>Título del Álbum</th>
                <th>Artista</th>
                <th>Año de Lanzamiento</th>
                <th>Género Musical</th>
              </tr>
            </thead>
            <tbody>
              {albums.map((album,index) => (
                <tr key={index}>
                  <td>{album.titulo_album}</td>
                  <td>{album.artista}</td>
                  <td>{album.anio_lanzamiento}</td>
                  <td>{album.genero_musical}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

export default App;


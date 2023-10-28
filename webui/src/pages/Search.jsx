import { Button, Container, Card, Row, Col, Form, ToggleButton } from 'react-bootstrap';
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import {Result} from '../components/Results';
// import "./styles.css";

const initialList = [
  { id: 0, title: 'Big Bellies', seen: false },
  { id: 1, title: 'Lunar Landscape', seen: false },
  { id: 2, title: 'Terracotta Army', seen: true },
];

function Search() {
  const [movies, setMovies] = useState([]);
  const [shows, setShows] = useState([]);
  const [episodes, setEpisodes] = useState([]);
  const [books, setBooks] = useState([]);
  const [games, setGames] = useState([]);
  const [people, setPeople] = useState([]);
  const [companies, setCompanies] = useState([]);
  const [groups, setGroups] = useState([]);
  const [query, setQuery] = useState([]);

  const [movieBool, setMovieBool] = useState(true);
  const [showBool, setShowBool] = useState(true);
  const [episodeBool, setEpisodeBool] = useState(true);
  const [bookBool, setBookBool] = useState(true);
  const [gameBool, setGameBool] = useState(true);
  const [peopleBool, setPeopleBool] = useState(true);
  const [companieBool, setCompanieBool] = useState(true);
  const [groupBool, setGroupBool] = useState(true);
  const [provBool, setProvBool] = useState(initialList);
  
  const [status, setStatus] = useState("");
  const [provs, setProvs] = useState([])


  function handleIncrementClick(index, val) {
    setProvBool(provBool.map(artwork => {
      if (artwork.id === index) {
        // Create a *new* object with changes
        return { ...artwork, seen: val };
      } else {
        // No changes
        return artwork;
      }
    }));
  }
  
  useEffect(() => {
    axios.get(import.meta.env.VITE_BACKEND + '/prov')
    .then(function (response) {
      // handle success
      // console.log(response.data)
      setProvs(response.data);
    })
    .catch(function (error) {
      // handle error
      setStatus(error);
      console.log(error);
    })
  }, [])

  async function searchApis(url) {
    
    axios.get(url + '/search?q=' + query + '&types=' + types)
    .then(function (response) {
      setStatus("Sucsess")
    })
    .catch(function (error) {
      setStatus(error);
      console.log(error);
    });
  }

  function getSearch() {

    axios.get(import.meta.env.VITE_BACKEND + '/media/search?q=' + query + '&types=' + types)
    .then(function (response) {
      // handle success
      // console.log(response.data)
      setMovies(response.data.movies);
      setShows(response.data.shows);
      setEpisodes(response.data.episodes);
      setBooks(response.data.books);
      setGames(response.data.Games);
      setPeople(response.data.people);
      setCompanies(response.data.companies);
      setGroups(response.data.groups);
      // setData(response.data);
    })
    .catch(function (error) {
      // handle error
      setStatus(error);
      console.log(error);
    })
  }

    return (

        <Container>
            <h1>Search </h1>
            <p>{status}</p>

            <div className="d-flex">
                <Form.Control
                    type="search"
                    placeholder="Search"
                    className="me-2"
                    onChange={(event) => setQuery(event.target.value)}
                    value={query}
                    onKeyPress={(event) => {if (event.charCode===13) {getSearch()}}}
                />
                <Button onClick={getSearch}>Search</Button>
            </div>

                <br />

                
                <ToggleButton
                  className="mb-2"
                  id='movieBool'
                  type="checkbox"
                  variant="outline-success"
                  checked={movieBool}
                  onChange={(e) => setMovieBool(e.currentTarget.checked)}
                >
                  Movie
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='showbool'
                  type="checkbox"
                  variant="outline-success"
                  checked={showBool}
                  onChange={(e) => setShowBool(e.currentTarget.checked)}
                >
                  Show
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='episodebool'
                  type="checkbox"
                  variant="outline-success"
                  checked={episodeBool}
                  onChange={(e) => setEpisodeBool(e.currentTarget.checked)}
                >
                  Episode
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='bookbool'
                  type="checkbox"
                  variant="outline-success"
                  checked={bookBool}
                  onChange={(e) => setBookBool(e.currentTarget.checked)}
                >
                  Book
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='gamebool'
                  type="checkbox"
                  variant="outline-success"
                  checked={gameBool}
                  onChange={(e) => setGameBool(e.currentTarget.checked)}
                >
                  Game
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='peoplebool'
                  type="checkbox"
                  variant="outline-success"
                  checked={peopleBool}
                  onChange={(e) => setPeopleBool(e.currentTarget.checked)}
                >
                  People
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='companybool'
                  type="checkbox"
                  variant="outline-success"
                  checked={companieBool}
                  onChange={(e) => setCompanieBool(e.currentTarget.checked)}
                >
                  Company
                </ToggleButton>
                <ToggleButton
                  className="mb-2"
                  id='groupbool'
                  type="checkbox"
                  variant="outline-success"
                  checked={groupBool}
                  onChange={(e) => setGroupBool(e.currentTarget.checked)}
                >
                  Collections
                </ToggleButton>

            <br />
            <div>
              {provBool.map(post => (
              <ToggleButton
                className="mb-2"
                id={post.id}
                type="checkbox"
                variant="outline-success"
                checked={post.seen}
                onChange={(e) => handleIncrementClick(post.id, e.currentTarget.checked)}
              >
                {post.title}
              </ToggleButton>

              ))}

            </div>
            <h2>provs</h2><Row>
            {provs.map(post => ( <p>{post.name}</p>))}
            </Row>

            <h2>Movies</h2><Row>
            {movies.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Shows</h2><Row>
            {shows.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Episodes</h2><Row>
            {episodes.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Books</h2><Row>
            {books.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Games</h2><Row>
            {games.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>People</h2><Row>
            {people.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Companies</h2><Row>
            {companies.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Collections</h2><Row>
            {groups.map(post => ( <Result post={post}/>))}
            </Row>
            

            

        </Container>

    );
}

export default Search;

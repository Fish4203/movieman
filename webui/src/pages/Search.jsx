import { Button, Container, Card, Row, Col, Form, ToggleButton, OverlayTrigger, Tooltip, InputGroup } from 'react-bootstrap';
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { MovieCard } from "../components/ResultsMovie";
import { ShowCard } from '../components/ResultsShow';
import { BookCard } from '../components/ResultsBook';
import { GameCard } from '../components/ResultsGame';
import { PersonCard } from '../components/ResultsPerson';
import { CompanyCard } from '../components/ResultsCompany';
import { GroupCard } from '../components/ResultsGroup';
import { EpisodeCard } from '../components/ResultsEpisode';
// import "./styles.css";


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
  const [provBool, setProvBool] = useState([]);
  
  const [status, setStatus] = useState("");


  function handleProvToggle(index, val) {
    setProvBool(provBool.map(provbooli => {
      if (provbooli.id === index) {
        return { ...provbooli, enabled: val };
      } else {
        // No changes
        return provbooli;
      }
    }));
  }
  
  useEffect(() => {
    axios.get(import.meta.env.VITE_BACKEND + '/prov')
    .then(function (response) {
      // handle success
      // console.log(response.data)
      let data = response.data;
      data.forEach(element => {
        element["enabled"] = true;
      });

      setProvBool(data);

    })
    .catch(function (error) {
      // handle error
      setStatus(error);
      console.log(error);
    })
  }, [])

  function getTypes() {
    let types = "";
    if (movieBool) {
      types += "m";
    }
    if (showBool) {
      types += "s";
    }
    if (episodeBool) {
      types += "e";
    }
    if (bookBool) {
      types += "b";
    }
    if (gameBool) {
      types += "v";
    }
    if (peopleBool) {
      types += "p";
    }
    if (companieBool) {
      types += "c";
    }
    if (groupBool) {
      types += "g";
    }
    return types;
  }


  async function searchApis() {
    let types = getTypes();

    let urls = [];
    provBool.map((prov) => {
      if (prov.enabled) {
        urls.push(prov.domain + '/search?q=' + query + '&types=' + types);
      }
    });

    const requests = urls.map((url) => axios.get(url));

    await axios.all(requests).then(axios.spread((...responses) => {
      responses.forEach((resp) => {
        // console.log("this should be first");
        console.log(resp.status);
      });
    }));

    console.log("why");
      
    getSearch(); 
    }

  function getSearch() {
    let types = getTypes();

    axios.get(import.meta.env.VITE_BACKEND + '/media/search?q=' + query + '&types=' + types)
    .then(function (response) {
      // handle success
      // console.log(response.data)
      if (response.data.movies != null) {
        setMovies(response.data.movies);
      } else {
        setMovies([]);
      }

      if (response.data.shows != null) {
        setShows(response.data.shows);
      } else {
        setShows([]);
      }
      if (response.data.episodes != null) {
        setEpisodes(response.data.episodes);
      } else {
        setEpisodes([]);
      }
      if (response.data.books != null) {
        setBooks(response.data.books);
      } else {
        setBooks([]);
      }
      if (response.data.Games != null) {
        setGames(response.data.Games);
      } else {
        setGames([]);
      }
      if (response.data.people != null) {
        setPeople(response.data.people);
      } else {
        setPeople([]);
      }
      if (response.data.companies != null) {
        setCompanies(response.data.companies);
      } else {
        setCompanies([]);
      }
      if (response.data.groups != null) {
        setGroups(response.data.groups);
      } else {
        setGroups([]);
      }
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
              <InputGroup className='m-2'>
                <Form.Control
                    type="search"
                    placeholder="Search"
                    onChange={(event) => setQuery(event.target.value)}
                    value={query}
                    onKeyDown={(event) => {if (event.charCode===13) {searchApis()}}}
                    onKeyPress={(event) => {if (event.charCode===13) {searchApis()}}}
                />
                <Button variant='success' onClick={searchApis}>Full Search</Button>
                <Button onClick={getSearch}>quick Search</Button>
              </InputGroup>
            </div>

            <Row>
              <Col>
                <Card className='m-2'>
                  <Card.Header>Search categories</Card.Header>
                  <Card.Body>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='movieBool'
                      type="checkbox"
                      variant="outline-info"
                      checked={movieBool}
                      onChange={(e) => setMovieBool(e.currentTarget.checked)}
                    >
                      Movie
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='showbool'
                      type="checkbox"
                      variant="outline-info"
                      checked={showBool}
                      onChange={(e) => setShowBool(e.currentTarget.checked)}
                    >
                      Show
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='episodebool'
                      type="checkbox"
                      variant="outline-info"
                      checked={episodeBool}
                      onChange={(e) => setEpisodeBool(e.currentTarget.checked)}
                    >
                      Episode
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='bookbool'
                      type="checkbox"
                      variant="outline-info"
                      checked={bookBool}
                      onChange={(e) => setBookBool(e.currentTarget.checked)}
                    >
                      Book
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='gamebool'
                      type="checkbox"
                      variant="outline-info"
                      checked={gameBool}
                      onChange={(e) => setGameBool(e.currentTarget.checked)}
                    >
                      Game
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='peoplebool'
                      type="checkbox"
                      variant="outline-info"
                      checked={peopleBool}
                      onChange={(e) => setPeopleBool(e.currentTarget.checked)}
                    >
                      People
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='companybool'
                      type="checkbox"
                      variant="outline-info"
                      checked={companieBool}
                      onChange={(e) => setCompanieBool(e.currentTarget.checked)}
                    >
                      Company
                    </ToggleButton>
                    <ToggleButton
                      className="mb-2 ms-1"
                      id='groupbool'
                      type="checkbox"
                      variant="outline-info"
                      checked={groupBool}
                      onChange={(e) => setGroupBool(e.currentTarget.checked)}
                    >
                      Collections
                    </ToggleButton>
                  </Card.Body>
                  <Card.Footer className="text-muted">Click to select/unselect to limit what is being searched for</Card.Footer>
                </Card>
              </Col>
              <Col>
                <Card className='m-2'>
                  <Card.Header>External databases</Card.Header>
                  <Card.Body>
                    {provBool.map(post => (
                      <OverlayTrigger
                      placement="right"
                      overlay={
                        <Tooltip id={`tooltip-right`}>
                          {post.description}
                        </Tooltip>
                      }
                    >
                      <ToggleButton
                        className="mb-2 ms-1"
                        id={post.id}
                        type="checkbox"
                        variant="outline-info"
                        checked={post.enabled}
                        onChange={(e) => handleProvToggle(post.id, e.currentTarget.checked)}
                      >
                        {post.name}
                      </ToggleButton>
                    </OverlayTrigger>

                    ))}
                  </Card.Body>
                  <Card.Footer className="text-muted">Click to select/unselect to pick the sites to search </Card.Footer>
                </Card>
              </Col>
            </Row>


            <br />

            <h2>Movies</h2>
            <Row>
            {movies.map(post => ( <MovieCard post={post}/>))}
            </Row>
            <h2>Shows</h2><Row>
            {shows.map(post => ( <ShowCard post={post}/>))}
            </Row>
            <h2>Episodes</h2><Row>
            {episodes.map(post => ( <EpisodeCard post={post}/>))}
            </Row>
            <h2>Books</h2><Row>
            {books.map(post => ( <BookCard post={post}/>))}
            </Row>
            <h2>Games</h2><Row>
            {games.map(post => ( <GameCard post={post}/>))}
            </Row>
            <h2>People</h2><Row>
            {people.map(post => ( <PersonCard post={post}/>))}
            </Row>
            <h2>Companies</h2><Row>
            {companies.map(post => ( <CompanyCard post={post}/>))}
            </Row>
            <h2>Collections</h2><Row>
            {groups.map(post => ( <GroupCard post={post}/>))}
            </Row>
            

            

        </Container>

    );
}

export default Search;

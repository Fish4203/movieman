import { Button, Container, Card, Row, Col, ToggleButton, OverlayTrigger, Tooltip } from 'react-bootstrap';
import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { MovieDetails } from "../components/ResultsMovie";
import { ShowDetails } from "../components/ResultsShow";
import { BookDetails } from "../components/ResultsBook";
import { GameDetails } from "../components/ResultsGame";
import { PersonDetails } from "../components/ResultsPerson";
import { CompanyDetails } from "../components/ResultsCompany";
import { GroupDetails } from "../components/ResultsGroup";
import { EpisodeDetails } from '../components/ResultsEpisode';



function Details() {
    const { type, id} = useParams();
    const [post, setPost] = useState([]);
    const [prov, setProv] = useState([]);
  
    const [status, setStatus] = useState("");


    function handleProvToggle(index, val) {
      setProv(prov.map(provi => {
        if (provi.id === index) {
          return { ...provi, enabled: val };
        } else {
          // No changes
          return provi;
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

        setProv(data);

      })
      .catch(function (error) {
        // handle error
        setStatus(error);
        console.log(error);
      })
    }, [])


    useEffect(() => {
        axios.get(import.meta.env.VITE_BACKEND + '/media/' + type + "/" + id)
        .then(function (response) {
          // handle success
          // console.log(response.data)
          setPost(response.data);
    
        })
        .catch(function (error) {
          // handle error
          console.log(error);
        })
      }, [])


      async function update() {
        let urls = [];
        let typeChar = "";
        let key = "";
        if (type == "movie") {
          typeChar = "m";
          key = "movies";
        } else if (type == "show") {
          typeChar = "s";
          key = "shows";
        } else if (type == "book") {
          typeChar = "b";
          key = "books";
        } else if (type == "game") {
          typeChar = "v";
          key = "games";
        } else if (type == "person") {
          typeChar = "p";
          key = "people";
        } else if (type == "company") {
          typeChar = "c";
          key = "companies";
        } else if (type == "group") {
          typeChar = "g";
          key = "group";
        }

        let extids = post[key][0].externalIds;

        console.log(prov);

        prov.map((provi) => {
          if (provi.enabled && provi.types.includes(typeChar)) {
            urls.push(provi.domain + '/' + type + '/' + extids[provi.name]);
          }
        });
    
        const requests = urls.map((url) => axios.get(url));
    
        await axios.all(requests).then(axios.spread((...responses) => {
          responses.forEach((resp) => {
            // console.log("this should be first");
            console.log(resp.status);
          });
        }));
    
        window.location.reload(false);
      }



    const renderDetails = () => {
      if ('people' in post) {
        return <PersonDetails post={post["people"][0]} movies={post.movies} shows={post.shows} books={post.books} games={post.games} />
      } else if ('companies' in post) {
        return <CompanyDetails post={post["companies"][0]} movies={post.movies} shows={post.shows} books={post.books} games={post.games} />
      } else if ('groups' in post) {
        return <GroupDetails post={post["groups"][0]} movies={post.movies} shows={post.shows} books={post.books} games={post.games} />
      } else if ('movies' in post) {
        return <MovieDetails post={post["movies"][0]}/>
      } else if ('shows' in post) {
        if ('seasons' in post) {
          return <ShowDetails post={post["shows"][0]} seasons={post["seasons"]} episodes={post["episodes"]} />
        }
        return <ShowDetails post={post["shows"][0]} />
      } else if ('episodes' in post) {
        return <EpisodeDetails post={post["episodes"][0]}/>
      } else if ('books' in post) {
        return <BookDetails post={post["books"][0]}/>
      } else if ('games' in post) {
        return <GameDetails post={post["games"][0]}/>
      } else {
        return <p>Error couldnt identify media type</p>
      }
    }

    return (
    <Container>
      <br />
      <Row>
        <Col xs={12} md={8}>
          {renderDetails()}
        </Col>
        <Col xs={12} md={4}>
        <Card>
          <Card.Header>
            <Card.Title>Update Details</Card.Title>
          </Card.Header>
          <Card.Body>
            <p>Click to select/unselect to pick the sites to search</p>
            {prov.map(post => (
              <OverlayTrigger
              placement="right"
              overlay={
                <Tooltip id={`tooltip-right`}>
                  {post.types}
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
            
            <br />

            <Button className='ms-1' onClick={update}>Update</Button>            
          </Card.Body>
        </Card>
        </Col>
      </Row>

        


    </Container>
 );
}

export default Details;

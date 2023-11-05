import { Button, Container, Card, Row, Col, ToggleButton, OverlayTrigger, Tooltip } from 'react-bootstrap';
import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { MovieDetails } from "../components/ResultsMovie";



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
        let typeChar = ""
        if (type == "movie") {
          typeChar = "m";
        }

        prov.map((provi) => {
          if (provi.enabled && provi.types.includes(typeChar)) {
            urls.push(provi.domain + '/' + type + '/' + post.externalIds["tmdb"]);
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
    

    return (
    <Container>
      <br />
      <Row>
        <Col xs={12} md={8}>
          {'movies' in post ? <MovieDetails post={post["movies"][0]}/>: ""}
        </Col>
        <Col xs={6} md={4}>
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

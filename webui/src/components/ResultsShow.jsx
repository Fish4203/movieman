import { Button, Container, Card, Row, Col, Carousel, Tabs, Tab, Badge, Stack } from 'react-bootstrap';
import { Outlet, Link, useNavigate } from "react-router-dom";
import "../assets/result.css";

function ShowCard(post) {
    const navigate = useNavigate();
    post = post.post;
    return (
    <Card onClick={() => navigate('/details/show/'+ post.id)} className='m-1 p-0' style={{ width: '18rem' }}>
        <Card.Img variant="top" src={'images' in post ? post.images[0] : "https://placehold.co/200x100"} />
        <Card.Body>
            <Card.Title>{'title' in post ? post.title : post.name}</Card.Title>
            <Card.Text>
                {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
            </Card.Text>
        </Card.Body>
        <Card.Footer>
            Release date: {post.date}
            {'length' in post ? <>Seasons: {post.seasons}</> : ""}
        </Card.Footer>
    </Card>
    );
}

function ShowDetails(args) {
    const navigate = useNavigate();
    const post = args.post;
    const seasons = args.seasons;
    const episodes = args.episodes;
    return (
    <div>
        <h2>{post.title}</h2>
        
        <Container >

        {'images' in post ? 
        <Carousel >
            {post.images.map(image => (
                <Carousel.Item >
                    <div className='centered'>
                    <img className='im'  src={image} alt="https://placehold.co/200x100" />

                    </div>
                </Carousel.Item>
            ))}
        </Carousel>
        : <p>No images found</p>}
        </Container>
        {'genre' in post ? <div>
            <Stack direction="horizontal" gap={2} className='m-3'>
                {post.genre.map(obj => (
                    <Badge pill bg="success">
                        {obj}
                    </Badge>
                ))}
            </Stack>
        </div>: ""}
        <h3>Description</h3>
        <p>{post.description}</p>
        <br />

        <Tabs
        defaultActiveKey="other"
        id="uncontrolled-tab-example"
        className="mb-3"
        >
            <Tab eventKey="other" title="Other facts">
                <p>Relese date: {post.date}</p>
                {'length' in post ? <p>Seasons: {post.seasons} min</p> : ""}
                {'info' in post ? <Button variant="info" href={post.info}>Info</Button>: ""}
            </Tab>
            {'reviews' in post ? <div>
                <Tab eventKey="Reviews" title="Reviews" >
                    <ul>
                        {Object.keys(post.reviews).map(obj =>
                            (<li>{obj}: {post.reviews[obj]}</li>)
                        )}
                    </ul>
                </Tab>
            </div>: <Tab eventKey="Reviews" title="Reviews" disabled></Tab>}

            {'platforms' in post ? <div>
                <Tab eventKey="Platforms" title="Platforms">
                    <ul>
                        {post.platforms.map(obj =>
                            (<li>{obj}</li>)
                        )}
                    </ul>
                </Tab>
            </div>: <Tab eventKey="Platforms" title="Platforms" disabled></Tab>}

            {'externalIds' in post ? 
                <Tab eventKey="externalIds" title="External Ids" >
                    {'externalIds' in post ? 
                    <ul>
                        {Object.keys(post.externalIds).map(obj =>
                            (<li>{obj}: {post.externalIds[obj]}</li>)
                        )}
                    </ul>
                    : <p>No external ids</p>}
                </Tab>
            : <Tab eventKey="externalIds" title="External Ids" disabled></Tab>}
        </Tabs>

        {seasons != null ? 
        <Tabs 
        id="tabsSeason"
        className="mb-3">
            {seasons.map(season => 
            <Tab eventKey={season.seasonId} title={season.seasonId}>
                <p>Relese Date: {season.date}</p>
                <p>Description: {season.description}</p>

                <ul>
                    {episodes.map(episode => (
                        <>{episode.seasonId == season.seasonId ? <li > <a onClick={() => {navigate('/details/episode/'+ post.id + "_" + episode.seasonId + '_' + episode.episodeId); navigate(0);}}>{episode.title}</a></li> : ""}</>
                    ))}
                </ul>
            </Tab>
            )}
        </Tabs>
        : ""}


    </div>
    );
}

export {ShowCard, ShowDetails};

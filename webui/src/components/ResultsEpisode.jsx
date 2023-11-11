import { Button, Container, Card, Row, Col, Carousel, Tabs, Tab, Badge, Stack } from 'react-bootstrap';
import { Outlet, Link, useNavigate } from "react-router-dom";
import "../assets/result.css";
import { useEffect } from 'react';

function EpisodeCard(post) {
    const navigate = useNavigate();
    post = post.post;
    return (
    <Card onClick={() => navigate('/details/episode/'+ post.showId + '_' + post.seasonId + '_' + post.episodeId)} className='m-1 p-0' style={{ width: '18rem' }}>
        <Card.Img variant="top" src={'images' in post ? post.images[0] : "https://placehold.co/200x100"} />
        <Card.Body>
            <Card.Title>{'title' in post ? post.title : post.name}</Card.Title>
            <Card.Text>
                {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
            </Card.Text>
        </Card.Body>
        <Card.Footer>
            Release date: {post.date}
            Season: {post.seasonId}
            Episode: {post.episodeId}
        </Card.Footer>
    </Card>
    );
}

function EpisodeDetails(args) {
    const navigate = useNavigate();
    const post = args.post;
    const season = args.season;
    const show = args.show;
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
        {'genre' in show ? <div>
            <Stack direction="horizontal" gap={2} className='m-3'>
                {show.genre.map(obj => (
                    <Badge pill bg="success">
                        {obj}
                    </Badge>
                ))}
            </Stack>
        </div>: ""}
        <h3>Description</h3>
        <p>{post.description}</p>
        <p>Relese date: {post.date}</p>
        <p>Episode: {post.episodeId } of {season.episodes}</p>

        {post.episodeId != 1 ?<Button variant="succsess" onClick={() => {navigate('/details/episode/'+ post.showId + '_' + post.seasonId + '_' + (post.episodeId-1)); navigate(0);}}>Previous Episode</Button> : ""}
        <Button variant="succsess" onClick={() => {navigate('/details/show/'+ show.id); navigate(0);}}>Show</Button>
        {post.episodeId != season.episodes ? <Button variant="succsess" onClick={() => {navigate('/details/episode/'+ post.showId + '_' + post.seasonId + '_' + (post.episodeId +1)); navigate(0);}}>Next Episode</Button>: ""}

        <br />

        {'reviews' in post ? <div>
                <ul>
                    {Object.keys(post.reviews).map(obj =>
                        (<li>{obj}: {post.reviews[obj]}</li>)
                    )}
                </ul>
        </div>: ""}

        


    </div>
    );
}

export {EpisodeCard, EpisodeDetails};

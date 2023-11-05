import { Button, Container, Card, Row, Col, Carousel } from 'react-bootstrap';
import { Outlet, Link, useNavigate } from "react-router-dom";
import "../assets/result.css";

function MovieCard(post) {
    const navigate = useNavigate();
    post = post.post;
    return (
    <Card onClick={() => navigate('/details/movie/'+ post.id)} className='m-1 p-0' style={{ width: '18rem' }}>
        <Card.Img variant="top" src={'images' in post ? post.images[0] : "https://placehold.co/200x100"} />
        <Card.Body>
            <Card.Title>{'title' in post ? post.title : post.name}</Card.Title>
            <Card.Text>
                {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
            </Card.Text>
        </Card.Body>
        <Card.Footer>
            Release date: {post.date}
            {'length' in post ? <>Length: {post.length}</> : ""}
        </Card.Footer>
    </Card>
    );
}

function MovieDetails(post) {
    const navigate = useNavigate();
    post = post.post;
    return (
    <div>
        <h2>{post.title}</h2>
        
        <Container >

        <Carousel >
            {post.images.map(image => (
                <Carousel.Item >
                    <div className='centered'>
                    <img className='im'  src={image} alt="https://placehold.co/200x100" />

                    </div>
                </Carousel.Item>
            ))}
        </Carousel>
        </Container>
        <br />
        <h3>Description</h3>
        <p>{post.description}</p>

        {'genre' in post ? <div>
            <h3>Genres</h3>
            <ul>
                {post.genres.map(obj =>
                    (<li>{obj}</li>)
                )}
            </ul>
        </div>: ""}

        <h3>Other facts</h3>
        <p>Relese date: {post.date}</p>
        {'budget' in post ? <p>Budget: {post.budget}</p> : ""}
        {'length' in post ? <p>Length: {post.length}</p> : ""}
        {'rating' in post ? <p>Age Rating: {post.rating}</p> : ""}
        {'info' in post ? <Button variant="info" href={post.info}>Info</Button>: ""}

        {'reviews' in post ? <div>
            <h3>Reviews</h3>
            <ul>
                {Object.keys(post.reviews).map(obj =>
                    (<li>{obj}: {post.reviews[obj]}</li>)
                )}
            </ul>
        </div>: ""}

        {'platforms' in post ? <div>
            <h3>Platforms</h3>
            <ul>
                {post.platforms.map(obj =>
                    (<li>{obj}</li>)
                )}
            </ul>
        </div>: ""}

        <p>External Ids:</p>
        <ul>
            {Object.keys(post.externalIds).map(obj =>
                (<li>{obj}: {post.externalIds[obj]}</li>)
            )}
        </ul>
    </div>
    );
}

export {MovieCard, MovieDetails};

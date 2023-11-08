import { Button, Container, Card, Row, Col, Carousel, Tabs, Tab, Badge, Stack } from 'react-bootstrap';
import { Outlet, Link, useNavigate } from "react-router-dom";
import "../assets/result.css";

function PersonCard(post) {
    const navigate = useNavigate();
    post = post.post;
    return (
    <Card onClick={() => navigate('/details/person/'+ post.id)} className='m-1 p-0' style={{ width: '18rem' }}>
        <Card.Img variant="top" src={'images' in post ? post.images[0] : "https://placehold.co/200x100"} />
        <Card.Body>
            <Card.Title>{'name' in post ? post.name : ""}</Card.Title>
            <Card.Text>
                {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
            </Card.Text>
        </Card.Body>
        <Card.Footer>
            Release date: {post.date}
            {'length' in post ? <>Title: {post.role}</> : ""}
        </Card.Footer>
    </Card>
    );
}

function PersonDetails(args) {
    const post = args.post;
    const movies = args.movies;
    const shows = args.shows;
    const books = args.books;
    const games = args.games;
    return (
    <div>
        <h2>{post.name}</h2>
        
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
        <p>Date of birth: {post.date}</p>
        <br />

        <Tabs
        defaultActiveKey="movies"
        id="uncontrolled-tab-example"
        className="mb-3"
        >
            <Tab eventKey="movies" title="Movies">
                {movies != null ? 
                <ul>
                    {movies.map(movie => (
                        <li>{movie.title}</li>
                    ))}
                </ul>
                : <p>No movies found</p>}
            </Tab>
            <Tab eventKey="shows" title="Shows">
                {shows != null ? 
                <ul>
                    {shows.map(show => (
                        <li>{show.title}</li>
                    ))}
                </ul>
                : <p>No shows found</p>}
            </Tab>
            <Tab eventKey="books" title="Books">
                {books != null ? 
                <ul>
                    {books.map(book => (
                        <li>{book.title}</li>
                    ))}
                </ul>
                : <p>No books found</p>}
            </Tab>
            <Tab eventKey="games" title="Games">
                {games != null ? 
                <ul>
                    {games.map(game => (
                        <li>{game.title}</li>
                    ))}
                </ul>
                : <p>No games found</p>}
                </Tab>

            {'externalIds' in post ? 
                <Tab eventKey="externalIds" title="External Ids" >
                    <ul>
                        {Object.keys(post.externalIds).map(obj =>
                            (<li>{obj}: {post.externalIds[obj]}</li>)
                        )}
                    </ul>
                </Tab>
            : <Tab eventKey="externalIds" title="External Ids" disabled></Tab>}
        </Tabs>


    </div>
    );
}

export {PersonCard, PersonDetails};

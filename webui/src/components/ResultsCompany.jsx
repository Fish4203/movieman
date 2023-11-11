import { Button, Container, Card, Row, Col, Carousel, Tabs, Tab, Badge, Stack } from 'react-bootstrap';
import { Outlet, Link, useNavigate } from "react-router-dom";
import "../assets/result.css";

function CompanyCard(post) {
    const navigate = useNavigate();
    post = post.post;
    return (
    <Card onClick={() => navigate('/details/company/'+ post.id)} className='m-1 p-0' style={{ width: '18rem' }}>
        <Card.Img variant="top" src={'images' in post ? post.images[0] : "https://placehold.co/200x100"} />
        <Card.Body>
            <Card.Title>{post.name}</Card.Title>
            <Card.Text>
                {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
            </Card.Text>
        </Card.Body>
        <Card.Footer>
            Founding date: {post.date}
        </Card.Footer>
    </Card>
    );
}

function CompanyDetails(args) {
    const navigate = useNavigate();
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
        <h3>Description</h3>
        <p>{post.description}</p>
        <p>Founding date: {post.date}</p>
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
                        <li ><a onClick={() => {navigate('/details/movie/'+ movie.id); navigate(0);}}>{movie.title}</a></li>
                    ))}
                </ul>
                : <p>No movies found</p>}
            </Tab>
            <Tab eventKey="shows" title="Shows">
                {shows != null ? 
                <ul>
                    {shows.map(show => (
                        <li onClick={() => {navigate('/details/show/'+ show.id); navigate(0);}}>{show.title}</li>
                    ))}
                </ul>
                : <p>No shows found</p>}
            </Tab>
            <Tab eventKey="books" title="Books">
                {books != null ? 
                <ul>
                    {books.map(book => (
                        <li onClick={() => {navigate('/details/book/'+ book.id); navigate(0);}}>{book.title}</li>
                    ))}
                </ul>
                : <p>No books found</p>}
            </Tab>
            <Tab eventKey="games" title="Games">
                {games != null ? 
                <ul>
                    {games.map(game => (
                        <li onClick={() => {navigate('/details/game/'+ game.id); navigate(0);}}>{game.title}</li>
                    ))}
                </ul>
                : <p>No games found</p>}
                </Tab>

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


    </div>
    );
}

export {CompanyCard, CompanyDetails};

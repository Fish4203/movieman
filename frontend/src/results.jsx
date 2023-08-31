import { Button, Container, Card, Row, Col, Stack } from 'react-bootstrap';
// import axios from 'axios';
// import "./styles.css";

function Movie({movies}) {
    if (movies == null) {
        return (<></>);
    } 
    return (
        <><h2>Movies</h2><Row>
            {movies.map(post => (
                <Card className='m-2' style={{ width: '18rem' }}>
                    <Card.Img variant="top" src="https://placehold.co/200x100" />
                    <Card.Body>
                        <Card.Title>{post.title}</Card.Title>
                        <Card.Text>
                            {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
                        </Card.Text>
                        <Button variant="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
            ))}
        </Row></>
    );
}

function People({people}) {
    if (people == null) {
        return (<></>);
    } 
    return (
        <><h2>People</h2><Row>
            {people.map(post => (
                <Card className='m-2' style={{ width: '18rem' }}>
                    <Card.Img variant="top" src="https://placehold.co/200x100" />
                    <Card.Body>
                        <Card.Title>{post.name}</Card.Title>
                        <Card.Text>
                            {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
                        </Card.Text>
                        <Button variant="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
            ))}
        </Row></>
    );
}

function Show({shows}) {
    if (shows == null) {
        return (<></>);
    } 
    return (
        <><h2>Shows</h2><Row>
            {shows.map(post => (
                <Card className='m-2' style={{ width: '18rem' }}>
                    <Card.Img variant="top" src="https://placehold.co/200x100" />
                    <Card.Body>
                        <Card.Title>{post.title}</Card.Title>
                        <Card.Text>
                            {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
                        </Card.Text>
                        <Button variant="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
            ))}
        </Row></>
    );
}

function Episode({episodes}) {
    if (episodes == null) {
        return (<></>);
    } 
    return (
        <><h2>Episodes</h2><Row>
            {episodes.map(post => (
                <Card className='m-2' style={{ width: '18rem' }}>
                    <Card.Img variant="top" src="https://placehold.co/200x100" />
                    <Card.Body>
                        <Card.Title>{post.title}</Card.Title>
                        <Card.Text>
                            {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
                        </Card.Text>
                        <Button variant="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
            ))}
        </Row></>
    );
}

export {Movie, Show, People, Episode};

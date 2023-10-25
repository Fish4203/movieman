import { Button, Container, Card, Row, Col, Stack } from 'react-bootstrap';
// import axios from 'axios';
// import "./styles.css";

function Movie({movies}) {
    if (movies == null) {
        return (<><h2>Movies</h2></>);
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
        return (<><h2>People</h2></>);
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

function Show({shows, lim}) {
    if (shows == null) {
        return (<><h2>Shows</h2></>);
    } 
    return (
        <><h2>Shows</h2><Row>
            {shows.slice(0, lim).map(post => (
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
        return (<><h2>Episodes</h2></>);
    } 
    return (
        <></>
    );
}

function Result(post) {
    post = post.post;
    console.log(post);
    return (
        <Card className='m-2' style={{ width: '18rem' }}>
        <Card.Img variant="top" src={'images' in post ? "images/" + post.images[0] : "images/"} />
        <Card.Body>
            <Card.Title>{'title' in post ? post.title : post.name}</Card.Title>
            <Card.Text>
                {post.description.length > 250 ? `${post.description.substring(0, 250)}...` : post.description}
            </Card.Text>
            {/* <Button variant="primary">Go somewhere</Button> */}
        </Card.Body>
    </Card>
    );

}

export {Result};

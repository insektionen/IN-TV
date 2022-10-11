import React, {useContext, useEffect, useState} from 'react';
import {Screen, Slideshow} from '../API/Types';
import {ApiContext} from '../App';
import {Button, Card, Col, Container, Dropdown, DropdownButton, Row} from 'react-bootstrap';
import {useNavigate} from 'react-router-dom';

const ScreenTile = (props: { screen: Screen, selected: boolean, onClick: (s: Screen) => void }) => {
    return (
        <Col style={{userSelect: 'none', cursor: 'pointer'}}>
            <Card border={props.selected ? 'primary' : ''} onClick={() => props.onClick(props.screen)}>
                <Card.Body>
                    <Card.Title style={{textAlign: 'center'}}>{props.screen.name}</Card.Title>
                </Card.Body>
            </Card>
        </Col>
    );
};

const ScreenSelector = (props: { screens: Screen[], onClick: (s: Screen) => void, selectedScreens: string[] }) => {
    return (
        <>
            <Row>
                <Col>
                    <h3>Select screens:</h3>
                </Col>
            </Row>
            {props.screens.length > 0 ?
                <Row xs={1} md={4} className={'g-4'}>
                    {
                        props.screens.map(screen => {
                            return <ScreenTile key={screen.name} screen={screen} onClick={props.onClick}
                                selected={props.selectedScreens.indexOf(screen.name) !== -1} />;
                        })
                    }
                </Row> : <Row><Col><h4>No screens available...</h4></Col></Row>}
        </>
    );
};

const SlideshowDropdown = (props: { slideshows: Slideshow[], onSelect: (sl: Slideshow) => void }) => {
    const [selected, setSelected] = useState<Slideshow>();
    const onSelect = (sl: Slideshow) => {
        setSelected(sl);
        props.onSelect(sl);
    };

    return (
        <Row>
            <Col>
                <h3>Select slideshow:</h3>
                <DropdownButton title={selected ? selected.name : 'Select...'}>
                    {
                        props.slideshows.map(sl => <Dropdown.Item key={sl.name}
                            onClick={() => onSelect(sl)}>{sl.name}</Dropdown.Item>)
                    }
                </DropdownButton>
            </Col>
        </Row>
    );
};

export const Start = () => {
    const [slideshows, setSlideshows] = useState<Slideshow[]>([]);
    const [selectedSlideshow, setSelectedSlideshow] = useState<Slideshow>();
    const [screens, setScreens] = useState<Screen[]>([]);
    const [selectedScreens, setSelectedScreens] = useState<Screen[]>([]);
    const navigate = useNavigate();
    const cli = useContext(ApiContext);

    useEffect(() => {
        cli.getSlideshows().then(res => setSlideshows(res));
        cli.getStatus().then(res => setScreens(res.connected_screens));
    }, []);

    const handleStart = () => {
        if (selectedSlideshow && selectedScreens.length > 0) {
            cli.startSlideshow(selectedSlideshow.name, selectedScreens).then(() => {
                navigate('/');
            });
        }
    };

    const onScreenClick = (s: Screen) => {
        const index = selectedScreens.indexOf(s);
        if (index === -1) {
            setSelectedScreens([...selectedScreens, s]);
        } else {
            const newData = [...selectedScreens];
            newData.splice(index, 1);
            setSelectedScreens(newData);
        }
    };

    return (
        <Container>
            <Row>
                <Col>
                    <h2>Start slideshow</h2>
                </Col>
            </Row>
            <SlideshowDropdown slideshows={slideshows} onSelect={setSelectedSlideshow} />
            <ScreenSelector screens={screens} onClick={onScreenClick}
                selectedScreens={selectedScreens.map((s) => s.name)} />
            <Row className={'my-2'}>
                <Col>
                    <Button variant={'success'} onClick={handleStart}
                        disabled={selectedSlideshow === undefined || selectedScreens.length === 0}>Start</Button>
                </Col>
            </Row>
        </Container>
    );
};
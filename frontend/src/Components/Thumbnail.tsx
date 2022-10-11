import {ReactChild} from 'react';
import {Button, Card, Col} from 'react-bootstrap';
import {Slide} from '../API/Types';

interface ThumbnailProps {
    name: string,
    children: ReactChild | ReactChild[],
    slide: Slide,
    onClick?: (s: Slide) => void,
    onEditClick?: () => void,
    onRemoveClick?: () => void
}

export const Thumbnail = (props: ThumbnailProps) => {
    const onClickHandler = () => {
        if (props.onClick) {
            props.onClick(props.slide);
        }
    };
    const cursor = props.onClick ? 'pointer' : 'move';

    return (
        <Col style={{maxWidth: '18rem', cursor: cursor, userSelect: 'none'}} onClick={onClickHandler}>
            <Card style={{minHeight: '20rem'}}>
                <Card.Body style={{
                    overflow: 'hidden',
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'space-between'
                }}>
                    <div style={{overflow: 'hidden'}}>
                        <Card.Title>{props.name}</Card.Title>
                        {props.children}
                    </div>
                    {props.slide.timeout > 0 ?
                        <Card.Text><i className={'bi-clock'} /> {props.slide.timeout}s</Card.Text> : null}
                </Card.Body>
                {(props.onEditClick || props.onRemoveClick) ?
                    <Card.Footer>
                        {props.onEditClick ?
                            <Button variant={'secondary'} onClick={props.onEditClick}><i
                                className={'bi-pencil'} /></Button> : null}
                        {props.onRemoveClick ?
                            <Button variant={'danger'} style={{float: 'right'}} onClick={props.onRemoveClick}><i
                                className={'bi-trash'} /></Button> : null}
                    </Card.Footer> : null
                }
            </Card>
        </Col>
    );
};
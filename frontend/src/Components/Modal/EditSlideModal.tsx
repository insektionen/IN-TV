import {Button, Modal} from 'react-bootstrap';
import {SlideEditor} from '../Slides/SlideEditor';
import React, {useState} from 'react';
import {Slide} from '../../API/Types';

interface EditSlideModalProps {
    visible: boolean,
    onClose: () => void,
    onSave: (s: Slide) => void,
    okButtonText: string,
    slide: Slide
}

export const EditSlideModal = (props: EditSlideModalProps) => {
    const [slideData, setSlideData] = useState(props.slide.data);
    const [timeout, setTimeout] = useState(props.slide.timeout);
    const onOkClick = () => {
        const newSlide = {...props.slide, timeout: timeout, data: slideData};
        props.onSave(newSlide);
    };

    return (
        <Modal show={props.visible} onHide={props.onClose}>
            <Modal.Header closeButton>
                <Modal.Title>Settings:</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <SlideEditor
                    slide={{position: props.slide.position, type: props.slide.type, data: slideData, timeout: timeout}}
                    setSlideData={setSlideData} setTimeout={setTimeout} />
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={props.onClose}>Close</Button>
                <Button variant="primary" onClick={onOkClick}>{props.okButtonText}</Button>
            </Modal.Footer>
        </Modal>);
};
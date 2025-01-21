import {Slide} from '../../API/Types';
import {PictureEditor} from './Picture/PictureEditor';
import {TextEditor} from './Text/TextEditor';
import {ClockEditor} from './Clock/ClockEditor';
import {ChangeEvent, ReactChild} from 'react';
import {Form} from 'react-bootstrap';

export interface SlideEditorProps {
    slide: Slide,
    // eslint-disable-next-line
    setSlideData: (data: Record<string, any>) => void
    setTimeout: (timeout: number) => void
}

export const BaseEditor = (props: { children: ReactChild | ReactChild[] } & SlideEditorProps) => {
    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        props.setTimeout(Number(event.target.value));
    };

    return (
        <Form>
            {props.children}
            <Form.Group controlId={'editor.Timeout'}>
                <Form.Label>Timeout (sec):</Form.Label>
                <Form.Control type={'text'} placeholder={'5'} value={props.slide.timeout > 0 ? props.slide.timeout : ''}
                    onChange={handleChange} />
            </Form.Group>
        </Form>
    );
};

export const SlideEditor = (props: SlideEditorProps) => {
    switch (props.slide.type) {
    case 'picture':
        return <PictureEditor {...props} />;
    case 'text':
        return <TextEditor {...props} />;
    case 'clock':
        return <ClockEditor {...props} />;
    default:
        return <h4>Unknown type</h4>;
    }
};
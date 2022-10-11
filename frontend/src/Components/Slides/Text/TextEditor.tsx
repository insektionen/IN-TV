import {Form} from 'react-bootstrap';
import {BaseEditor, SlideEditorProps} from '../SlideEditor';
import {ChangeEvent} from 'react';

export const TextEditor = (props: SlideEditorProps) => {
    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        const name = event.target.name;
        const newData = {...props.slide.data};
        newData[name] = event.target.value;
        props.setSlideData(newData);
    };
    return (
        <BaseEditor {...props}>
            <Form.Group controlId={'text_edit.title'}>
                <Form.Label>Title:</Form.Label>
                <Form.Control type={'text'} placeholder={'Title text'} value={props.slide.data['title'] ?? ''}
                    name={'title'}
                    onChange={handleChange} />
            </Form.Group>
            <Form.Group controlId={'text_edit.text'}>
                <Form.Label>Text:</Form.Label>
                <Form.Control as={'textarea'} placeholder={'Some text...'} rows={4}
                    value={props.slide.data['text'] ?? ''}
                    name={'text'} onChange={handleChange} />
            </Form.Group>
        </BaseEditor>
    );
};
import {Form} from 'react-bootstrap';
import {BaseEditor, SlideEditorProps} from '../SlideEditor';
import {ChangeEvent} from 'react';

export const ClockEditor = (props: SlideEditorProps) => {
    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const newData = {...props.slide.data};
        newData[e.target.name] = e.target.checked;
        props.setSlideData(newData);
    };

    return (
        <BaseEditor {...props}>
            <Form.Check type={'checkbox'} label={'With seconds'} onChange={handleChange}
                value={props.slide.data['with_seconds'] ?? false} name={'with_seconds'} />
            <Form.Check type={'checkbox'} label={'With date'} onChange={handleChange}
                value={props.slide.data['with_date'] ?? false} name={'with_date'} />
        </BaseEditor>
    );
};
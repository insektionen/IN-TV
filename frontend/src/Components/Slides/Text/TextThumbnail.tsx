import {Slide} from '../../../API/Types';
import {Thumbnail} from '../../Thumbnail';

export const TextThumbnail = (props: { slide: Slide, onClick?: (s: Slide) => void, onRemoveClick?: () => void, onEditClick?: () => void }) => {
    return (
        <Thumbnail name={'Text'} slide={props.slide} onClick={props.onClick} onRemoveClick={props.onRemoveClick}
            onEditClick={props.onEditClick}>
            <p style={{fontWeight: 'bold'}}>{props.slide.data['title'] ?? 'Title text'}</p>
            <p>{props.slide.data['text'] ?? 'Text'}</p>
        </Thumbnail>
    );
};
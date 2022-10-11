import {Slide} from '../../../API/Types';
import {Thumbnail} from '../../Thumbnail';

export const PictureThumbnail = (props: { slide: Slide, onClick?: (s: Slide) => void, onRemoveClick?: () => void, onEditClick?: () => void }) => {
    const url = props.slide.data['url'] ?? 'img/example_picture.png';
    return (
        <Thumbnail name={'Picture'} slide={props.slide} onClick={props.onClick} onRemoveClick={props.onRemoveClick}
            onEditClick={props.onEditClick}>
            <img src={url} alt="Picture preview" width={'100%'} />
        </Thumbnail>
    );
};
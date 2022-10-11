import {Slide} from '../../API/Types';
import {PictureThumbnail} from './Picture/PictureThumbnail';
import {TextThumbnail} from './Text/TextThumbnail';
import {ClockThumbnail} from './Clock/ClockThumbnail';

export const SlideThumbnail = (props: { slide: Slide, onRemove?: (s: Slide) => void, onEdit?: (s: Slide) => void }) => {
    const onRemoveChecked = () => {
        if (props.onRemove) {
            props.onRemove(props.slide);
        }
    };
    const onEditChecked = () => {
        if (props.onEdit) {
            props.onEdit(props.slide);
        }
    };
    switch (props.slide.type) {
    case 'picture':
        return <PictureThumbnail slide={props.slide}
            onRemoveClick={props.onRemove ? onRemoveChecked : undefined}
            onEditClick={props.onEdit ? onEditChecked : undefined} />;
    case 'text':
        return <TextThumbnail slide={props.slide}
            onRemoveClick={props.onRemove ? onRemoveChecked : undefined}
            onEditClick={props.onEdit ? onEditChecked : undefined} />;
    case 'clock':
        return <ClockThumbnail slide={props.slide}
            onRemoveClick={props.onRemove ? onRemoveChecked : undefined}
            onEditClick={props.onEdit ? onEditChecked : undefined} />;
    default:
        return <h4>Unknown slide type</h4>;
    }
};
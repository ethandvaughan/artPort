import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

const ImageCarousel = ({ images }) => {
  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
  };

  return (
    <div className='w-64'>
      <Slider {...settings}>
        {images.map((image) => (
          <div key={image}>
            <img src={image} />
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default ImageCarousel;

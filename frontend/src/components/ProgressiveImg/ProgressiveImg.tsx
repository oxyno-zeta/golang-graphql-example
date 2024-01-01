import React, { ComponentProps, useEffect, useState } from 'react';

type Props = {
  placeholderSrc: string;
  src: string;
  alt: string;
} & ComponentProps<'img'>;

// Inspired from: https://blog.logrocket.com/progressive-image-loading-react-tutorial/
function ProgressiveImg({ placeholderSrc, src, alt, ...props }: Props) {
  const [imgSrc, setImgSrc] = useState(placeholderSrc || src);

  useEffect(() => {
    const img = new Image();
    img.src = src;
    img.onload = () => {
      setImgSrc(src);
    };
  }, [src]);

  return <img src={imgSrc} alt={alt} {...props} />;
}

export default ProgressiveImg;

import React, { ComponentProps, useEffect, useState } from 'react';

type Props = {
  readonly placeholderSrc: string;
  readonly src: string;
  readonly alt: string;
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

  return <img alt={alt} src={imgSrc} {...props} />;
}

export default ProgressiveImg;

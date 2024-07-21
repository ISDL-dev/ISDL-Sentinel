import { chakra, forwardRef, HTMLChakraProps, ThemingProps } from '@chakra-ui/react';
import { mode } from '@chakra-ui/theme-tools';

// `Card` コンポーネントのスタイル設定を含む
const baseStyle = (props: any) => ({
  p: '20px',
  display: 'flex',
  flexDirection: 'column',
  width: '100%',
  position: 'relative',
  borderRadius: '20px',
  alignItems: 'center',
  minWidth: '0px',
  wordWrap: 'break-word',
  bg: mode('#ffffff', 'navy.800')(props),
  backgroundClip: 'border-box',
  height: { base: 500, md: 600 }
});

interface CustomCardProps extends HTMLChakraProps<'div'>, ThemingProps {}

const CustomCard = forwardRef<CustomCardProps, 'div'>((props, ref) => {
  const { size, variant, ...rest } = props;
  const styles = baseStyle(props);

  return <chakra.div ref={ref} __css={styles} {...rest} />;
});

CustomCard.defaultProps = {
  size: 'md',
  variant: 'default'
};

export default CustomCard;

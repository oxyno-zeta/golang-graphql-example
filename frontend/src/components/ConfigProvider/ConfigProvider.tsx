import React, { useEffect, useState, ReactNode } from 'react';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { useTranslation } from 'react-i18next';
import axios from 'axios';
import ConfigContext from '../../contexts/ConfigContext';
import { ConfigModel } from '../../models/config';

interface Props {
  children: ReactNode;
  loadingComponent: ReactNode;
  skipConfigLoader?: boolean;
}

const defaultProps = { skipConfigLoader: false };

function ConfigProvider({ children, loadingComponent, skipConfigLoader }: Props) {
  // Check if skip is enabled
  if (skipConfigLoader) {
    return <>{children}</>;
  }

  // Create loading state
  const [loading, setLoading] = useState(true);
  // Create error raised state
  const [errorRaised, setErrorRaised] = useState(false);
  // Configuration
  const [config, setConfig] = useState<ConfigModel | null>(null);

  const { t } = useTranslation();

  useEffect(() => {
    axios
      .get('/config/config.json')
      .then(({ data: cfg }) => {
        setLoading(false);
        setConfig(cfg);
      })
      .catch((err) => {
        setLoading(false);
        setErrorRaised(true);
        console.error(err);
      });
  }, []);

  // Check if loading
  if (loading) {
    return <>{loadingComponent}</>;
  }

  // Check if error have been raised
  if (errorRaised) {
    return (
      <Box
        sx={{
          alignItems: 'center',
          left: '50%',
          top: '50%',
          transform: 'translate(-50%, -50%)',
          position: 'absolute',
        }}
      >
        <Typography color="error">{t('common.configLoadError')}</Typography>
      </Box>
    );
  }

  // Check if configuration have been loaded
  if (!config) {
    return (
      <Box
        sx={{
          alignItems: 'center',
          left: '50%',
          top: '50%',
          transform: 'translate(-50%, -50%)',
          position: 'absolute',
        }}
      >
        <Typography color="error">{t('common.configEmptyError')}</Typography>
      </Box>
    );
  }

  return <ConfigContext.Provider value={config}>{children}</ConfigContext.Provider>;
}

ConfigProvider.defaultProps = defaultProps;

export default ConfigProvider;

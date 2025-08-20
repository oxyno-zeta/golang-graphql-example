import React, { useCallback, useEffect, useMemo, useState, memo, useRef } from 'react';
import Grid from '@mui/material/Grid';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import ButtonGroup from '@mui/material/ButtonGroup';
import Box from '@mui/material/Box';
import { mdiPlus, mdiDelete, mdiPlusBoxMultiple } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';
import { type FilterDefinitionFieldsModel } from '../../../../../models/general';
import FilterBuilderField from '../FilterBuilderField';
import { generateKey, buildFieldInitialValue, buildFilterBuilderInitialItems } from '../../utils';
import {
  type LineOrGroup,
  type BuilderInitialValueObject,
  type FieldInitialValueObject,
  type FilterValueObject,
} from '../../types';

/* eslint-disable @typescript-eslint/no-explicit-any */

export interface Props {
  readonly filterDefinitionModel: FilterDefinitionFieldsModel;
  readonly initialValue: BuilderInitialValueObject;
  readonly acceptEmptyLines?: boolean;
  readonly onRemove?: () => void | undefined;
  readonly onChange: (fo: null | FilterValueObject) => void;
}

function FilterBuilder({
  filterDefinitionModel,
  onRemove = undefined,
  onChange,
  initialValue,
  acceptEmptyLines = false,
}: Props) {
  // Setup translate
  const { t } = useTranslation();
  // States
  const resultsRef = useRef<Record<string, null | any>>({});
  const [groupKey, setGroupKey] = useState(initialValue.group);
  const [items, setItems] = useState<LineOrGroup[]>(initialValue.items);

  // Watch initialValue
  useEffect(() => {
    // Set initial value data
    setGroupKey(initialValue.group);
    setItems(initialValue.items);
    // Clean result reference
    const resRefKeys = Object.keys(resultsRef.current);
    resRefKeys.forEach((key) => {
      if (initialValue.items.findIndex((it) => it.key === key) === -1) {
        // Key isn't found => Need to clean it
        delete resultsRef.current[key];
      }
    });
  }, [initialValue]);

  const localSaveManagementHandler = useCallback(
    (newGroupKey: string) => {
      // Create new items
      const newItems = Object.values(resultsRef.current);

      // Check if array is empty and if the accept empty lines isn't enabled
      if (newItems.length === 0 && !acceptEmptyLines) {
        // Don't accepting this
        onChange(null);
        return;
      }

      // Check if array is empty and if the accept empty lines is enabled
      if (newItems.length === 0 && acceptEmptyLines) {
        // Accepting this
        onChange({});
        return;
      }

      // Find one item that doesn't have value
      const withoutValueIndex = newItems.findIndex((v) => !v);
      // Check if there is an item without value
      if (withoutValueIndex !== -1) {
        onChange(null);
        return;
      }

      // Optimization if there is only 1 value
      // Return directly the value
      if (newItems.length === 1) {
        onChange(newItems[0]);
        return;
      }

      // Otherwise create and save object
      onChange({ [newGroupKey]: newItems });
    },
    [acceptEmptyLines, onChange],
  );

  // Add group handler
  const addGroupHandler = useCallback(() => {
    setItems((v) => {
      // Generate key
      const key = generateKey('group');

      // Create new items list
      const newItems = [
        ...v,
        {
          type: 'group',
          key,
          initialValue: buildFilterBuilderInitialItems(undefined),
        },
      ];
      // Save new result value
      resultsRef.current[key] = null;
      // Save management
      localSaveManagementHandler(groupKey);

      return newItems;
    });
  }, [groupKey, localSaveManagementHandler]);

  // Add line handler
  const addLineHandler = useCallback(() => {
    setItems((v) => {
      // Generate key
      const key = generateKey('line');

      // Create new items list
      const newItems = [...v, { type: 'line', key, initialValue: buildFieldInitialValue(undefined)[0] }];
      // Save new result value
      resultsRef.current[key] = null;
      // Save management
      localSaveManagementHandler(groupKey);

      return newItems;
    });
  }, [groupKey, localSaveManagementHandler]);

  // Local remove handler
  const localRemoveHandler = useCallback(
    (key: string) => () => {
      setItems((v) => {
        const newItems = v.filter((it) => it.key !== key);
        // Delete key in result ref
        delete resultsRef.current[key];
        // Save management
        localSaveManagementHandler(groupKey);

        return newItems;
      });
    },
    [groupKey, localSaveManagementHandler],
  );

  // Create all memorized save handlers
  const saveHandlers = useMemo(
    () =>
      items.map((it) => (v: any) => {
        // Save value
        resultsRef.current[it.key] = v;

        // Save management
        localSaveManagementHandler(groupKey);
      }),
    [items, groupKey, localSaveManagementHandler],
  );

  return (
    <Box sx={{ display: 'flex' }}>
      <Box
        sx={{
          borderLeftColor: 'text.secondary',
          borderLeftStyle: 'solid',
          borderLeftWidth: '1px',
          margin: '16px 10px 50px 0',
        }}
      />
      <Box sx={{ display: 'block', width: '100%' }}>
        <ButtonGroup color={items.length === 0 && !acceptEmptyLines ? 'error' : 'primary'} size="small">
          <Button
            onClick={() => {
              setGroupKey('AND');
              // Save management
              localSaveManagementHandler('AND');
            }}
            variant={groupKey === 'AND' ? 'contained' : undefined}
          >
            {t('common.operations.and')}
          </Button>
          <Button
            onClick={() => {
              setGroupKey('OR');
              // Save management
              localSaveManagementHandler('OR');
            }}
            variant={groupKey === 'OR' ? 'contained' : undefined}
          >
            {t('common.operations.or')}
          </Button>
        </ButtonGroup>
        <Tooltip title={<>{t('common.filter.addNewField')}</>}>
          <IconButton onClick={addLineHandler} sx={{ margin: '0 5px' }}>
            <SvgIcon>
              <path d={mdiPlus} />
            </SvgIcon>
          </IconButton>
        </Tooltip>
        <Tooltip title={<>{t('common.filter.addNewGroupField')}</>}>
          <IconButton onClick={addGroupHandler} sx={{ margin: '0 5px' }}>
            <SvgIcon>
              <path d={mdiPlusBoxMultiple} />
            </SvgIcon>
          </IconButton>
        </Tooltip>
        {onRemove ? (
          <Tooltip title={<>{t('common.filter.deleteGroupField')}</>}>
            <IconButton onClick={onRemove}>
              <SvgIcon>
                <path d={mdiDelete} />
              </SvgIcon>
            </IconButton>
          </Tooltip>
        ) : null}

        {items.map((it, index) => {
          if (it.type === 'line') {
            return (
              <Box data-testid={it.key} key={it.key} sx={{ display: 'flex', margin: '10px 0 5px 0' }}>
                <Box sx={{ display: 'flex', alignItems: 'center', margin: '-20px 5px 0 0' }}>
                  <Tooltip title={<>{t('common.filter.deleteField')}</>}>
                    <IconButton onClick={localRemoveHandler(it.key)}>
                      <SvgIcon>
                        <path d={mdiDelete} />
                      </SvgIcon>
                    </IconButton>
                  </Tooltip>
                </Box>
                <Grid
                  container
                  spacing={1}
                  sx={{
                    // Force line height to always include space for errors
                    minHeight: '72px',
                    flexGrow: 1,
                  }}
                >
                  <FilterBuilderField
                    filterDefinitionModel={filterDefinitionModel}
                    id={it.key}
                    initialValue={it.initialValue as FieldInitialValueObject}
                    onChange={saveHandlers[index]}
                  />
                </Grid>
              </Box>
            );
          }

          return (
            <FilterBuilder
              filterDefinitionModel={filterDefinitionModel}
              initialValue={it.initialValue as BuilderInitialValueObject}
              key={it.key}
              onChange={saveHandlers[index]}
              onRemove={localRemoveHandler(it.key)}
            />
          );
        })}
      </Box>
    </Box>
  );
}

export default memo(FilterBuilder);

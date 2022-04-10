import React, { memo } from 'react';
import { SortOrderModel, SortOrderFieldModel } from '../../../../models/general';
import SortField from '../SortField';

interface Props {
  sort: null | undefined | Record<string, SortOrderModel>;
  sortFields: SortOrderFieldModel[];
  onChange: (field: string, v: SortOrderModel) => void;
}

function SortForm({ sort, sortFields, onChange }: Props) {
  return (
    <div style={{ padding: '15px 10px 10px 10px' }}>
      {sortFields.map((fieldDeclaration, index) => (
        <div key={fieldDeclaration.field} style={{ marginBottom: sortFields.length - 1 !== index ? '15px' : '' }}>
          <SortField
            value={sort ? sort[fieldDeclaration.field] : undefined}
            fieldDeclaration={fieldDeclaration}
            onChange={(v) => {
              onChange(fieldDeclaration.field, v);
            }}
          />
        </div>
      ))}
    </div>
  );
}

export default memo(SortForm);

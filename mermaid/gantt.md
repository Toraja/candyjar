# Mermaid

## My Gantt Chart Template

### Config
```mermaid
gantt
    title Gantt Chart
    dateFormat  YYYY-MM-DD
    axisFormat  %m-%d(%a)
    tickInterval 1day
    excludes weekends
    %% ---
    %% <tasks goes here>
```

### Theme
#### Forest

```mermaid
%%{
  init: {
    'theme': 'forest',
    'themeVariables': {
      'background': '#111111',
      'fontSize': '22px',
      'critBkgColor': "#ee9922",
      'critBorderColor': "#bb7722"
    }
  }
}%%
gantt
    ...
    todayMarker stroke-width:3px,stroke:#13540c,opacity:0.5
```

#### Base
```mermaid
%%{
  init: {
    'theme': 'base',
    'themeVariables': {
      'background': '#111111',
      'primaryTextColor': '#222222',
      "titleColor": "#009955",
      'weekendBkgColor': "#ff0000",
      'taskBkgColor': "#00bb55",
      'taskBorderColor': "#009955",
      'activeTaskBkgColor': "#44aaee",
      'activeTaskBorderColor': "#2277ee",
      'critBkgColor': "#ffaa33",
      'critBorderColor': "#bb7722"
    }
  }
}%%
gantt
    ...
    todayMarker stroke-width:3px,stroke:#2277ee,opacity:0.5
```

## Theme variables
Other Gantt chart specific variables include:
- doneTaskBkgColor
- doneTaskBorderColor
- sectionBkgColor
- sectionBkgColor2
- altSectionBkgColor

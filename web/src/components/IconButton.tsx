export function IconButton({
  icon,
  size = "30",
  aClass = "",
  onClick = undefined,
}) {
  const classes = `IconButton ${aClass}`;
  return (
    <button class={classes} onClick={onClick}>
      <img src={icon} width={size} height={size} />
    </button>
  );
}

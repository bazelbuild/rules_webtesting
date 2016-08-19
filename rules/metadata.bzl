# A library of functions for working with web_test metadata files.


def merge_files(ctx, merger, output, inputs):
  """Takes a list of input metadata files, and produces a merged output file.
    params:
      ctx: a skylark rule context
      merger: the WTL metadata merger executable
      output: a File object for the output file
      inputs: a list of File objects for the input files
  """
  paths = [i.path for i in inputs]
  short_paths = [i.short_path for i in inputs]
  args = ['--output', output.path] + paths

  ctx.action(outputs=[output],
             inputs=inputs,
             executable=merger,
             arguments=args,
             mnemonic='METADATAMERGER',
             progress_message='merging %s' % (', '.join(short_paths)))


def create_file(ctx,
                output,
                capabilities=None,
                form_factor=None,
                browser_name=None,
                environment=None,
                browser_label=None,
                test_label=None,
                crop_screenshots=None,
                record_video=None):
  """Generates a web_test metadata file with specified contents."""
  content = '{\n  "_comment": "generated file for %s"' % ctx.label
  if capabilities:
    content += ',\n  "capabilities": ' + capabilities
  if form_factor:
    content += ',\n  "formFactor": "' + form_factor + '"'
  if browser_name:
    content += ',\n  "browserName": "' + browser_name + '"'
  if environment:
    content += ',\n  "environment": "' + environment + '"'
  if browser_label:
    content += ',\n  "browserLabel": "%s"' % browser_label
  if test_label:
    content += ',\n  "testLabel": "%s"' % test_label
  if crop_screenshots == True:
    content += ',\n  "cropScreenshots": true'
  elif crop_screenshots == False:
    content += ',\n  "cropScreenshots": false'
  if record_video:
    content += ',\n  "recordVideo": "' + record_video + '"'
  content += '\n}\n'

  ctx.file_action(output=output, content=content, executable=False)
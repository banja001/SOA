using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.Execution;
using Explorer.Tours.Core.Domain.Sessions;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Execution
{
    [Authorize(Policy = "touristPolicy")]
    [Route("api/tourist/session")]
    public class SessionController : BaseApiController
    {
        private readonly ISessionService _sessionService;
        private readonly IHttpClientFactory _factory;

        public SessionController(ISessionService sessionService, IHttpClientFactory factory)
        {
            _sessionService = sessionService;
            _factory = factory;
        }

        [HttpGet("{id:long}")]
        public ActionResult<SessionDto> Get(int id)
        {
            var result = _sessionService.Get(id);
            return CreateResponse(result);
        }

        [HttpGet("getByTouristId/{id:long}")]
        public ActionResult<SessionDto> GetActiveByTouristId(long id)
        {
            var result = _sessionService.GetActiveByTouristId(id);
            return CreateResponse(result);
        }

        [HttpGet("getAllByTouristId/{id:long}")]
        public ActionResult<SessionDto> GetAllByTouristId(long id)
        {
            var result = _sessionService.GetAllByTouristId(id);
            return CreateResponse(result);
        }
        [HttpGet("geActiveSessiontByTouristId/{id:long}")]
        public ActionResult<SessionDto> GetActiveSessionByTouristId(long id)
        {
            var result = _sessionService.GetActiveSessionByTouristId(id);
            return CreateResponse(result);
        }

        [HttpPost]
        public async Task<ActionResult<SessionDto>> Create([FromBody] SessionDto session)
        {
            // var result = _sessionService.Create(session);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PostAsJsonAsync("/sessions/create", session);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();
            var createdSession = System.Text.Json.JsonSerializer.Deserialize<SessionDto>(jsonResponse);

            return Ok(createdSession);
        }


        [HttpPut]
        public async Task<ActionResult<SessionDto>> Update([FromBody] SessionDto session)
        {
            //var result = _sessionService.Update(session);
            //return CreateResponse(result);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("/sessions/update", session);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();
            var updatedSession = System.Text.Json.JsonSerializer.Deserialize<SessionDto>(jsonResponse);

            return Ok(updatedSession);
        }

        [HttpGet("check/{id:long}")]
        public ActionResult<bool> Check(int id)
        {
            var result = _sessionService.ValidForTouristComment(id);
            return CreateResponse(result);
        }

        [HttpPut("completeKeyPoint/{sessionId:int}/{keyPointId:int}")]
        public async Task<ActionResult<SessionDto>> CompleteKeyPoint(int sessionId, int keyPointId)
        {
            //var result = _sessionService.AddCompletedKeyPoint(sessionId, keyPointId);
            //return CreateResponse(result);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("/sessions/completeKeypoint/" +  sessionId, keyPointId);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();
            var updatedSession = System.Text.Json.JsonSerializer.Deserialize<SessionDto>(jsonResponse);

            return Ok(updatedSession);
        }

        [HttpGet("getByTourAndTouristId/{tourId:long}/{touristId:long}")]
        public ActionResult<SessionDto> GetByTourAndTouristId(long tourId, long touristId)
        {
            var result = _sessionService.GetByTourAndTouristId(tourId, touristId);
            return CreateResponse(result);
        }

        [HttpGet("getSessionsByTouristId/{touristId:long}")]
        public ActionResult<SessionDto> GetPagedByTouristId(long touristId, [FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _sessionService.GetPagedByTouristId(touristId, page, pageSize);
            return CreateResponse(result);
        }
    }
}
